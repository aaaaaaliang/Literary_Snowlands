package service

import (
	"Literary_Snowlands/internal/model"
	"context"
)

type IUser interface {
	NewCaptcha(ctx context.Context) (id string, b64s string, answer string, err error)
	VerifyCaptcha(ctx context.Context, id string, inputCaptcha string) (bool, error)
	Create(ctx context.Context, in model.UserCreateInput) (token string, uid int, username string, err error)
	IsUsernameRepeat(ctx context.Context, in model.UserCreateInput) (bool, error)
	Login(ctx context.Context, in model.UserLoginInput) (token string, uid int, nickname string, err error)
}

var localUser IUser

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
