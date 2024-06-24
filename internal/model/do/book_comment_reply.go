// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BookCommentReply is the golang structure of table book_comment_reply for DAO operations like Where/Data.
type BookCommentReply struct {
	g.Meta       `orm:"table:book_comment_reply, do:true"`
	Id           interface{} // 主键
	CommentId    interface{} // 评论ID
	UserId       interface{} // 回复用户ID
	ReplyContent interface{} // 回复内容
	AuditStatus  interface{} // 审核状态;0-待审核 1-审核通过 2-审核不通过
	CreateTime   *gtime.Time // 创建时间
	UpdateTime   *gtime.Time // 更新时间
}
