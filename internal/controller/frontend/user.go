package frontend

import (
	"Literary_Snowlands/api/frontend"
	v1 "Literary_Snowlands/api/frontend/v1"
	"Literary_Snowlands/internal/model"
	"Literary_Snowlands/internal/service"
	"context"
)

type UserControllerV1 struct {
}

func NewUserControllerV1() frontend.IUserV1 {
	return &UserControllerV1{}
}

func (c *UserControllerV1) GetCaptcha(ctx context.Context, req *v1.GetCaptchaReq) (res *v1.GetCaptchaRes, err error) {
	sessionId, b64s, _, err := service.User().NewCaptcha(ctx)
	if err != nil {
		return nil, err
	} else {
		return &v1.GetCaptchaRes{Img: b64s, SessionId: sessionId}, nil
	}
}

func (c *UserControllerV1) SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error) {
	token, uid, nickname, err := service.User().Login(ctx, model.UserLoginInput{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.SignInRes{
		UId:      uid,
		Token:    token,
		Nickname: nickname,
	}, nil
}

func (c *UserControllerV1) SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error) {
	token, uid, username, err := service.User().Create(ctx, model.UserCreateInput{
		Username:  req.Username,
		Password:  req.Password,
		VelCode:   req.VelCode,
		SessionId: req.SessionId,
	})
	if err != nil {
		return nil, err
	}
	return &v1.SignUpRes{
		UId:      uid,
		Token:    token,
		Username: username,
	}, err
}
