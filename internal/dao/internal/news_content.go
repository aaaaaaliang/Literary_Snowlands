// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NewsContentDao is the data access object for table news_content.
type NewsContentDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns NewsContentColumns // columns contains all the column names of Table for convenient usage.
}

// NewsContentColumns defines and stores column names for table news_content.
type NewsContentColumns struct {
	Id         string // 主键
	NewsId     string // 新闻ID
	Content    string // 新闻内容
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// newsContentColumns holds the columns for table news_content.
var newsContentColumns = NewsContentColumns{
	Id:         "id",
	NewsId:     "news_id",
	Content:    "content",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewNewsContentDao creates and returns a new DAO object for table data access.
func NewNewsContentDao() *NewsContentDao {
	return &NewsContentDao{
		group:   "default",
		table:   "news_content",
		columns: newsContentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NewsContentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NewsContentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NewsContentDao) Columns() NewsContentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NewsContentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NewsContentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NewsContentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
