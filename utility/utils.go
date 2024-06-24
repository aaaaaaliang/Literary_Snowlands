package utility

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"golang.org/x/crypto/bcrypt"
	"os"
	"path/filepath"
	"time"
)

//func GenerateJWT() (string, error) {
//	token := jwt.New(jwt.SigningMethodHS256)
//	claims := token.Claims.(jwt.MapClaims)
//
//	claims["authorized"] = true
//	claims["user"] = "aliang"
//	claims["exp"] = time.Now().Add(120 * time.Minute).Unix()
//	// 获取环境变量中的密钥字符串，并转换为[]byte类型
//	JwtSecretKey := os.Getenv("JWT_SECRET")
//	signedString, err := token.SignedString([]byte(JwtSecretKey))
//	if err != nil {
//		return "", fmt.Errorf("生成签名出错 %v", err)
//	}
//	return signedString, nil
//}

func GenerateJWT(userId int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["userId"] = userId // 将用户ID加入到token
	claims["exp"] = time.Now().Add(120 * time.Minute).Unix()

	JwtSecretKey := os.Getenv("JWT_SECRET")
	signedString, err := token.SignedString([]byte(JwtSecretKey))
	if err != nil {
		return "", fmt.Errorf("生成签名出错 %v", err)
	}
	return signedString, nil
}

//func ParseJWT(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
//	JwtSecretKey := os.Getenv("JWT_SECRET")
//	fmt.Println("....", JwtSecretKey)
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//		}
//		return []byte(JwtSecretKey), nil
//	})
//	if err != nil {
//		return nil, nil, err
//	}
//
//	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//		return token, claims, nil
//	} else {
//		return nil, nil, fmt.Errorf("invalid token")
//	}
//}

//func ParseJWT(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
//	JwtSecretKey := os.Getenv("JWT_SECRET")
//	if JwtSecretKey == "" {
//		// 如果 JWT_SECRET 环境变量为空，则立即返回错误
//		return nil, nil, fmt.Errorf("JWT_SECRET environment variable is not set")
//	}
//
//	// 日志记录要解析的令牌字符串
//	fmt.Println("Parsing JWT:", tokenString)
//
//	// 解析 JWT 令牌
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		// 验证签名算法是否为预期的算法
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//		}
//		return []byte(JwtSecretKey), nil
//	})
//
//	// 如果在解析过程中出现错误，记录错误并返回
//	if err != nil {
//		fmt.Println("Error parsing JWT:", err)
//		return nil, nil, err
//	}
//
//	// 验证令牌是否有效
//	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//		// 令牌有效，返回解析后的令牌和声明
//		return token, claims, nil
//	} else {
//		// 令牌无效或声明类型不正确，返回错误
//		fmt.Println("Invalid or expired token")
//		return nil, nil, fmt.Errorf("invalid or expired token")
//	}
//}

// ParseJWT parses the JWT token string and returns a token object along with its claims
func ParseJWT(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	JwtSecretKey := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JwtSecretKey), nil
	})
	if err != nil {
		return nil, nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token, claims, nil
	} else {
		return nil, nil, fmt.Errorf("invalid token")
	}
}

func GenerateSalt(len int) (string, error) {
	bytes := make([]byte, len)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func HashPassword(password, salt string) (string, error) {
	passwordWithSalt := password + salt
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwordWithSalt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPasswordHash(inputPassword string, dbSalt string, hashPassword string) (bool, error) {
	passwordWithSalt := inputPassword + dbSalt
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(passwordWithSalt))
	if err != nil {
		return false, err
	}
	// 密码匹配
	return true, nil
}

func ResourceImages(ctx context.Context, file *ghttp.UploadFile) (url string, err error) {
	// 打开文件
	openFile, err := file.Open()
	if err != nil {
		return "", err
	}
	defer openFile.Close()

	//// 验证图片尺寸
	//img, _, err := image.DecodeConfig(openFile)
	//if err != nil {
	//	return "", fmt.Errorf("failed to decode image: %v", err)
	//}
	//if img.Width > consts.Max_Width || img.Height > consts.Max_Height {
	//	return "", fmt.Errorf("image dimensions (width: %d, height: %d) exceed the maximum allowed size", img.Width, img.Height)
	//}
	//if _, err := openFile.Seek(0, 0); err != nil {
	//	return "", fmt.Errorf("failed to reset file pointer: %v", err)
	//}

	// 从配置中获取七牛云的相关信息
	bucket := g.Cfg().MustGet(ctx, "qiniu.bucket").String()
	accessKey := g.Cfg().MustGet(ctx, "qiniu.accessKey").String()
	secretKey := g.Cfg().MustGet(ctx, "qiniu.secretKey").String()
	domain := g.Cfg().MustGet(ctx, "qiniu.url").String()

	// 初始化七牛云的认证信息
	mac := qbox.NewMac(accessKey, secretKey)

	// 设置上传策略
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	upToken := putPolicy.UploadToken(mac)

	// 配置上传参数
	cfg := storage.Config{
		Zone:     &storage.ZoneHuadong,
		UseHTTPS: false,
	}

	// 构建表单上传对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 生成文件在七牛云的存储路径（key）
	uniqueKey := uuid.New().String()
	ext := filepath.Ext(file.Filename)
	key := uniqueKey + ext

	// 获取文件大小
	fileHeader := file.FileHeader
	fileSize := fileHeader.Size

	putExtra := storage.PutExtra{}

	// 上传文件
	err = formUploader.Put(ctx, &ret, upToken, key, openFile, fileSize, &putExtra)
	if err != nil {
		return "", err
	}
	// 生成文件访问URL
	url = fmt.Sprintf("%s/%s", domain, ret.Key)

	return url, nil
}
