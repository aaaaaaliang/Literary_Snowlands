// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HomeBookDao is the data access object for table home_book.
type HomeBookDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns HomeBookColumns // columns contains all the column names of Table for convenient usage.
}

// HomeBookColumns defines and stores column names for table home_book.
type HomeBookColumns struct {
	Id         string //
	Type       string // 推荐类型;0-轮播图 1-顶部栏 2-本周强推 3-热门推荐 4-精品推荐
	Sort       string // 推荐排序
	BookId     string // 推荐小说ID
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// homeBookColumns holds the columns for table home_book.
var homeBookColumns = HomeBookColumns{
	Id:         "id",
	Type:       "type",
	Sort:       "sort",
	BookId:     "book_id",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewHomeBookDao creates and returns a new DAO object for table data access.
func NewHomeBookDao() *HomeBookDao {
	return &HomeBookDao{
		group:   "default",
		table:   "home_book",
		columns: homeBookColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HomeBookDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HomeBookDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HomeBookDao) Columns() HomeBookColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HomeBookDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HomeBookDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HomeBookDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
