// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserComment is the golang structure for table user_comment.
type UserComment struct {
	Id             uint64      `json:"id"             ` // 主键
	UserId         uint64      `json:"userId"         ` // 评论用户ID
	BookId         uint64      `json:"bookId"         ` // 评论小说ID
	CommentContent string      `json:"commentContent" ` // 评价内容
	ReplyCount     uint        `json:"replyCount"     ` // 回复数量
	AuditStatus    uint        `json:"auditStatus"    ` // 审核状态;0-待审核 1-审核通过 2-审核不通过
	CreateTime     *gtime.Time `json:"createTime"     ` // 创建时间
	UpdateTime     *gtime.Time `json:"updateTime"     ` // 更新时间
}
