// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BookContentDao is the data access object for table book_content.
type BookContentDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns BookContentColumns // columns contains all the column names of Table for convenient usage.
}

// BookContentColumns defines and stores column names for table book_content.
type BookContentColumns struct {
	Id         string // 主键
	ChapterId  string // 章节ID
	Content    string // 小说章节内容
	CreateTime string //
	UpdateTime string //
}

// bookContentColumns holds the columns for table book_content.
var bookContentColumns = BookContentColumns{
	Id:         "id",
	ChapterId:  "chapter_id",
	Content:    "content",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewBookContentDao creates and returns a new DAO object for table data access.
func NewBookContentDao() *BookContentDao {
	return &BookContentDao{
		group:   "default",
		table:   "book_content",
		columns: bookContentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BookContentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BookContentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BookContentDao) Columns() BookContentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BookContentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BookContentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BookContentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
