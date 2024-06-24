// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysLogDao is the data access object for table sys_log.
type SysLogDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns SysLogColumns // columns contains all the column names of Table for convenient usage.
}

// SysLogColumns defines and stores column names for table sys_log.
type SysLogColumns struct {
	Id         string //
	UserId     string // 用户id
	Username   string // 用户名
	Operation  string // 用户操作
	Time       string // 响应时间
	Method     string // 请求方法
	Params     string // 请求参数
	Ip         string // IP地址
	CreateTime string // 创建时间
}

// sysLogColumns holds the columns for table sys_log.
var sysLogColumns = SysLogColumns{
	Id:         "id",
	UserId:     "user_id",
	Username:   "username",
	Operation:  "operation",
	Time:       "time",
	Method:     "method",
	Params:     "params",
	Ip:         "ip",
	CreateTime: "create_time",
}

// NewSysLogDao creates and returns a new DAO object for table data access.
func NewSysLogDao() *SysLogDao {
	return &SysLogDao{
		group:   "default",
		table:   "sys_log",
		columns: sysLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysLogDao) Columns() SysLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
