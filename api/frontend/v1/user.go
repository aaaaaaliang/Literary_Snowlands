package v1

import "github.com/gogf/gf/v2/frame/g"

type GetCaptchaReq struct {
	g.Meta `path:"/api/front/resource/img_verify_code" method:"get" tags:"用户模块"  summary:"验证码"`
}

type GetCaptchaRes struct {
	Img       string `json:"img"`
	SessionId string `json:"sessionId"`
}

type SignInReq struct {
	g.Meta   `path:"/api/front/user/login" method:"post" tags:"用户模块"  summary:"用户登陆"`
	Username string `v:"required|length:6,16" json:"username"`
	Password string `v:"required|length:6,16" json:"password"`
}

type SignInRes struct {
	UId      int    `json:"uid"`
	Token    string `json:"token"`
	Nickname string `json:"nickName"`
}

type SignUpReq struct {
	g.Meta    `path:"/api/front/user/register" method:"post" tags:"用户模块"  summary:"用户注册"`
	Username  string `v:"required|length:6,16" json:"username"`
	Password  string `v:"required|length:6,16" json:"password"`
	VelCode   string `v:"required|length:4,4" json:"velCode"`
	SessionId string `v:"required" json:"sessionId"`
}

type SignUpRes struct {
	UId      int    `json:"uid"`
	Token    string `json:"token"`
	Username string `json:"username"`
}
