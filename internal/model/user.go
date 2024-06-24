package model

type UserCreateInput struct {
	Username  string
	Password  string
	VelCode   string
	SessionId string
}

type UserLoginInput struct {
	Username string
	Password string
}
