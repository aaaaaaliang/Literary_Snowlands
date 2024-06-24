// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthorInfo is the golang structure of table author_info for DAO operations like Where/Data.
type AuthorInfo struct {
	g.Meta        `orm:"table:author_info, do:true"`
	Id            interface{} // 主键
	UserId        interface{} // 用户ID
	InviteCode    interface{} // 邀请码
	PenName       interface{} // 笔名
	TelPhone      interface{} // 手机号码
	ChatAccount   interface{} // QQ或微信账号
	Email         interface{} // 电子邮箱
	WorkDirection interface{} // 作品方向;0-男频 1-女频
	Status        interface{} // 0：正常;1-封禁
	CreateTime    *gtime.Time // 创建时间
	UpdateTime    *gtime.Time // 更新时间
}
