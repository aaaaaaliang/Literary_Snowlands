// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserCommentReply is the golang structure for table user_comment_reply.
type UserCommentReply struct {
	Id           uint64      `json:"id"           ` // 主键
	CommentId    uint64      `json:"commentId"    ` // 评论ID
	UserId       uint64      `json:"userId"       ` // 回复用户ID
	ReplyContent string      `json:"replyContent" ` // 回复内容
	AuditStatus  uint        `json:"auditStatus"  ` // 审核状态;0-待审核 1-审核通过 2-审核不通过
	CreateTime   *gtime.Time `json:"createTime"   ` // 创建时间
	UpdateTime   *gtime.Time `json:"updateTime"   ` // 更新时间
}
