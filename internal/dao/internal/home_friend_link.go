// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HomeFriendLinkDao is the data access object for table home_friend_link.
type HomeFriendLinkDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns HomeFriendLinkColumns // columns contains all the column names of Table for convenient usage.
}

// HomeFriendLinkColumns defines and stores column names for table home_friend_link.
type HomeFriendLinkColumns struct {
	Id         string //
	LinkName   string // 链接名
	LinkUrl    string // 链接url
	Sort       string // 排序号
	IsOpen     string // 是否开启;0-不开启 1-开启
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// homeFriendLinkColumns holds the columns for table home_friend_link.
var homeFriendLinkColumns = HomeFriendLinkColumns{
	Id:         "id",
	LinkName:   "link_name",
	LinkUrl:    "link_url",
	Sort:       "sort",
	IsOpen:     "is_open",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewHomeFriendLinkDao creates and returns a new DAO object for table data access.
func NewHomeFriendLinkDao() *HomeFriendLinkDao {
	return &HomeFriendLinkDao{
		group:   "default",
		table:   "home_friend_link",
		columns: homeFriendLinkColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HomeFriendLinkDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HomeFriendLinkDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HomeFriendLinkDao) Columns() HomeFriendLinkColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HomeFriendLinkDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HomeFriendLinkDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HomeFriendLinkDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
