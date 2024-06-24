// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserInfo is the golang structure for table user_info.
type UserInfo struct {
	Id             uint64      `json:"id"             ` //
	Username       string      `json:"username"       ` // 登录名
	Password       string      `json:"password"       ` // 登录密码-加密
	Salt           string      `json:"salt"           ` // 加密盐值
	NickName       string      `json:"nickName"       ` // 昵称
	UserPhoto      string      `json:"userPhoto"      ` // 用户头像
	UserSex        uint        `json:"userSex"        ` // 用户性别;0-男 1-女
	AccountBalance uint64      `json:"accountBalance" ` // 账户余额
	Status         uint        `json:"status"         ` // 用户状态;0-正常
	CreateTime     *gtime.Time `json:"createTime"     ` // 创建时间
	UpdateTime     *gtime.Time `json:"updateTime"     ` // 更新时间
}
