// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HomeFriendLink is the golang structure of table home_friend_link for DAO operations like Where/Data.
type HomeFriendLink struct {
	g.Meta     `orm:"table:home_friend_link, do:true"`
	Id         interface{} //
	LinkName   interface{} // 链接名
	LinkUrl    interface{} // 链接url
	Sort       interface{} // 排序号
	IsOpen     interface{} // 是否开启;0-不开启 1-开启
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
}
