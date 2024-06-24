// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BookCategoryDao is the data access object for table book_category.
type BookCategoryDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns BookCategoryColumns // columns contains all the column names of Table for convenient usage.
}

// BookCategoryColumns defines and stores column names for table book_category.
type BookCategoryColumns struct {
	Id            string //
	WorkDirection string // 作品方向;0-男频 1-女频
	Name          string // 类别名
	Sort          string // 排序
	CreateTime    string // 创建时间
	UpdateTime    string // 更新时间
}

// bookCategoryColumns holds the columns for table book_category.
var bookCategoryColumns = BookCategoryColumns{
	Id:            "id",
	WorkDirection: "work_direction",
	Name:          "name",
	Sort:          "sort",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewBookCategoryDao creates and returns a new DAO object for table data access.
func NewBookCategoryDao() *BookCategoryDao {
	return &BookCategoryDao{
		group:   "default",
		table:   "book_category",
		columns: bookCategoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BookCategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BookCategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BookCategoryDao) Columns() BookCategoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BookCategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BookCategoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BookCategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
