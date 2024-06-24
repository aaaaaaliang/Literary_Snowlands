// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserFeedback is the golang structure for table user_feedback.
type UserFeedback struct {
	Id         uint64      `json:"id"         ` //
	UserId     uint64      `json:"userId"     ` // 反馈用户id
	Content    string      `json:"content"    ` // 反馈内容
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
}
