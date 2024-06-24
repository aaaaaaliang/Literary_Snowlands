// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NewsInfoDao is the data access object for table news_info.
type NewsInfoDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns NewsInfoColumns // columns contains all the column names of Table for convenient usage.
}

// NewsInfoColumns defines and stores column names for table news_info.
type NewsInfoColumns struct {
	Id           string // 主键
	CategoryId   string // 类别ID
	CategoryName string // 类别名
	SourceName   string // 新闻来源
	Title        string // 新闻标题
	CreateTime   string // 创建时间
	UpdateTime   string // 更新时间
}

// newsInfoColumns holds the columns for table news_info.
var newsInfoColumns = NewsInfoColumns{
	Id:           "id",
	CategoryId:   "category_id",
	CategoryName: "category_name",
	SourceName:   "source_name",
	Title:        "title",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
}

// NewNewsInfoDao creates and returns a new DAO object for table data access.
func NewNewsInfoDao() *NewsInfoDao {
	return &NewsInfoDao{
		group:   "default",
		table:   "news_info",
		columns: newsInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NewsInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NewsInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NewsInfoDao) Columns() NewsInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NewsInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NewsInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NewsInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
