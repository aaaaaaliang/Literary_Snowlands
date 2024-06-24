// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserCommentReplyDao is the data access object for table user_comment_reply.
type UserCommentReplyDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns UserCommentReplyColumns // columns contains all the column names of Table for convenient usage.
}

// UserCommentReplyColumns defines and stores column names for table user_comment_reply.
type UserCommentReplyColumns struct {
	Id           string // 主键
	CommentId    string // 评论ID
	UserId       string // 回复用户ID
	ReplyContent string // 回复内容
	AuditStatus  string // 审核状态;0-待审核 1-审核通过 2-审核不通过
	CreateTime   string // 创建时间
	UpdateTime   string // 更新时间
}

// userCommentReplyColumns holds the columns for table user_comment_reply.
var userCommentReplyColumns = UserCommentReplyColumns{
	Id:           "id",
	CommentId:    "comment_id",
	UserId:       "user_id",
	ReplyContent: "reply_content",
	AuditStatus:  "audit_status",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
}

// NewUserCommentReplyDao creates and returns a new DAO object for table data access.
func NewUserCommentReplyDao() *UserCommentReplyDao {
	return &UserCommentReplyDao{
		group:   "default",
		table:   "user_comment_reply",
		columns: userCommentReplyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserCommentReplyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserCommentReplyDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserCommentReplyDao) Columns() UserCommentReplyColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserCommentReplyDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserCommentReplyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserCommentReplyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
