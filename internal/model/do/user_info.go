// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserInfo is the golang structure of table user_info for DAO operations like Where/Data.
type UserInfo struct {
	g.Meta         `orm:"table:user_info, do:true"`
	Id             interface{} //
	Username       interface{} // 登录名
	Password       interface{} // 登录密码-加密
	Salt           interface{} // 加密盐值
	NickName       interface{} // 昵称
	UserPhoto      interface{} // 用户头像
	UserSex        interface{} // 用户性别;0-男 1-女
	AccountBalance interface{} // 账户余额
	Status         interface{} // 用户状态;0-正常
	CreateTime     *gtime.Time // 创建时间
	UpdateTime     *gtime.Time // 更新时间
}
