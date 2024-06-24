// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserFeedback is the golang structure of table user_feedback for DAO operations like Where/Data.
type UserFeedback struct {
	g.Meta     `orm:"table:user_feedback, do:true"`
	Id         interface{} //
	UserId     interface{} // 反馈用户id
	Content    interface{} // 反馈内容
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
}
