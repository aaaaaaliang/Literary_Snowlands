// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserBookshelfDao is the data access object for table user_bookshelf.
type UserBookshelfDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns UserBookshelfColumns // columns contains all the column names of Table for convenient usage.
}

// UserBookshelfColumns defines and stores column names for table user_bookshelf.
type UserBookshelfColumns struct {
	Id           string // 主键
	UserId       string // 用户ID
	BookId       string // 小说ID
	PreContentId string // 上一次阅读的章节内容表ID
	CreateTime   string // 创建时间;
	UpdateTime   string // 更新时间;
}

// userBookshelfColumns holds the columns for table user_bookshelf.
var userBookshelfColumns = UserBookshelfColumns{
	Id:           "id",
	UserId:       "user_id",
	BookId:       "book_id",
	PreContentId: "pre_content_id",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
}

// NewUserBookshelfDao creates and returns a new DAO object for table data access.
func NewUserBookshelfDao() *UserBookshelfDao {
	return &UserBookshelfDao{
		group:   "default",
		table:   "user_bookshelf",
		columns: userBookshelfColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserBookshelfDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserBookshelfDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserBookshelfDao) Columns() UserBookshelfColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserBookshelfDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserBookshelfDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserBookshelfDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
