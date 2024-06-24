// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserComment is the golang structure of table user_comment for DAO operations like Where/Data.
type UserComment struct {
	g.Meta         `orm:"table:user_comment, do:true"`
	Id             interface{} // 主键
	UserId         interface{} // 评论用户ID
	BookId         interface{} // 评论小说ID
	CommentContent interface{} // 评价内容
	ReplyCount     interface{} // 回复数量
	AuditStatus    interface{} // 审核状态;0-待审核 1-审核通过 2-审核不通过
	CreateTime     *gtime.Time // 创建时间
	UpdateTime     *gtime.Time // 更新时间
}
