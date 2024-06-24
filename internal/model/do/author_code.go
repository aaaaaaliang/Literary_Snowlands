// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthorCode is the golang structure of table author_code for DAO operations like Where/Data.
type AuthorCode struct {
	g.Meta       `orm:"table:author_code, do:true"`
	Id           interface{} // 主键
	InviteCode   interface{} // 邀请码
	ValidityTime *gtime.Time // 有效时间
	IsUsed       interface{} // 是否使用过;0-未使用 1-使用过
	CreateTime   *gtime.Time // 创建时间
	UpdateTime   *gtime.Time // 更新时间
}
