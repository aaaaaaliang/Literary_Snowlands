// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BookCommentCopy1Dao is the data access object for table book_comment_copy1.
type BookCommentCopy1Dao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns BookCommentCopy1Columns // columns contains all the column names of Table for convenient usage.
}

// BookCommentCopy1Columns defines and stores column names for table book_comment_copy1.
type BookCommentCopy1Columns struct {
	Id             string // 主键
	BookId         string // 评论小说ID
	UserId         string // 评论用户ID
	CommentContent string // 评价内容
	ReplyCount     string // 回复数量
	AuditStatus    string // 审核状态;0-待审核 1-审核通过 2-审核不通过
	CreateTime     string // 创建时间
	UpdateTime     string // 更新时间
}

// bookCommentCopy1Columns holds the columns for table book_comment_copy1.
var bookCommentCopy1Columns = BookCommentCopy1Columns{
	Id:             "id",
	BookId:         "book_id",
	UserId:         "user_id",
	CommentContent: "comment_content",
	ReplyCount:     "reply_count",
	AuditStatus:    "audit_status",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
}

// NewBookCommentCopy1Dao creates and returns a new DAO object for table data access.
func NewBookCommentCopy1Dao() *BookCommentCopy1Dao {
	return &BookCommentCopy1Dao{
		group:   "default",
		table:   "book_comment_copy1",
		columns: bookCommentCopy1Columns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BookCommentCopy1Dao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BookCommentCopy1Dao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BookCommentCopy1Dao) Columns() BookCommentCopy1Columns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BookCommentCopy1Dao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BookCommentCopy1Dao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BookCommentCopy1Dao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
