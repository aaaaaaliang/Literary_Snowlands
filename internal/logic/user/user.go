package book

import (
	"Literary_Snowlands/internal/dao"
	"Literary_Snowlands/internal/model"
	"Literary_Snowlands/internal/model/do"
	"Literary_Snowlands/internal/model/entity"
	"Literary_Snowlands/internal/service"
	"Literary_Snowlands/utility"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mojocn/base64Captcha"
)

type sUser struct {
}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

var store = base64Captcha.DefaultMemStore

// NewCaptcha 生成验证码
func (s *sUser) NewCaptcha(ctx context.Context) (id string, b64s string, answer string, err error) {
	//noiseCount: 背景噪点数量    showLineOptions:干扰线数量   length:验证码长度
	driverString := base64Captcha.NewDriverString(60, 100, 200, 2, 4, "123456abcdefghijklmnopqrstuvwxyz", nil, nil, nil)
	id, b64s, answer, err = base64Captcha.NewCaptcha(driverString, store).Generate()
	if err != nil {
		return "", "", "", err
	}
	ttl := int64(5 * 60)

	_, err = g.Redis().Set(ctx, "captcha:"+id, answer, gredis.SetOption{
		TTLOption: gredis.TTLOption{EX: &ttl},
		NX:        true,
	})
	if err != nil {
		return "", "", "", err
	}
	return id, b64s, answer, nil
}

func (s *sUser) VerifyCaptcha(ctx context.Context, id string, inputCaptcha string) (bool, error) {
	realAnswerVar, err := g.Redis().Get(ctx, "captcha:"+id)
	if err != nil {
		return false, err
	}
	realAnswer := realAnswerVar.String()
	return realAnswer == inputCaptcha, nil
}

func (s *sUser) Create(ctx context.Context, in model.UserCreateInput) (token string, uid int, username string, err error) {

	IsRepeat, err := s.IsUsernameRepeat(ctx, model.UserCreateInput{Username: in.Username})
	if err != nil {
		return "", 0, "", err
	}
	if IsRepeat {
		return "", 0, "", fmt.Errorf("用户名已存在")
	}

	isValid, err := s.VerifyCaptcha(ctx, in.SessionId, in.VelCode)
	if err != nil {
		return "", 0, "", err
	}
	if !isValid {
		return "", 0, "", fmt.Errorf("验证码错误")
	}

	salt, err := utility.GenerateSalt(2)
	if err != nil {
		return "", 0, "", err
	}

	hashPassword, err := utility.HashPassword(in.Password, salt)
	if err != nil {
		return "", 0, "", err
	}
	err = dao.UserInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.UserInfo.Ctx(ctx).Data(do.UserInfo{
			Username:       in.Username,
			Password:       hashPassword,
			Salt:           salt,
			NickName:       in.Username,
			UserPhoto:      nil,
			UserSex:        nil,
			AccountBalance: nil,
			Status:         0,
			CreateTime:     gtime.Now(),
			UpdateTime:     gtime.Now(),
		}).Insert()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", 0, "", err
	}

	var info entity.UserInfo
	err = dao.UserInfo.Ctx(ctx).Where("username = ?", in.Username).Scan(&info)
	if err != nil {
		return "", 0, "", err
	}

	username = info.Username
	uid = int(info.Id)
	token, err = utility.GenerateJWT(uid)
	if err != nil {
		return "", 0, "", err
	}

	return token, uid, username, nil
}

func (s *sUser) IsUsernameRepeat(ctx context.Context, in model.UserCreateInput) (bool, error) {
	count, err := dao.UserInfo.Ctx(ctx).Where("username = ?", in.Username).Count()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (s *sUser) Login(ctx context.Context, in model.UserLoginInput) (token string, uid int, nickname string, err error) {
	var userInfo entity.UserInfo
	err = dao.UserInfo.Ctx(ctx).Where("username = ?", in.Username).Scan(&userInfo)
	if err != nil {
		return "", 0, "", fmt.Errorf("该用户不存在，请前往注册")
	}
	isPasswordEqual, err := utility.CheckPasswordHash(in.Password, userInfo.Salt, userInfo.Password)
	if err != nil {
		return "", 0, "", err
	}
	if isPasswordEqual == false {
		return "", 0, "", fmt.Errorf("密码输入错误，请确认密码是否正确")
	}

	token, err = utility.GenerateJWT(int(userInfo.Id))
	if err != nil {
		return "", 0, "", err
	}
	return token, int(userInfo.Id), userInfo.NickName, nil
}
