// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserFeedbackDao is the data access object for table user_feedback.
type UserFeedbackDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns UserFeedbackColumns // columns contains all the column names of Table for convenient usage.
}

// UserFeedbackColumns defines and stores column names for table user_feedback.
type UserFeedbackColumns struct {
	Id         string //
	UserId     string // 反馈用户id
	Content    string // 反馈内容
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// userFeedbackColumns holds the columns for table user_feedback.
var userFeedbackColumns = UserFeedbackColumns{
	Id:         "id",
	UserId:     "user_id",
	Content:    "content",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewUserFeedbackDao creates and returns a new DAO object for table data access.
func NewUserFeedbackDao() *UserFeedbackDao {
	return &UserFeedbackDao{
		group:   "default",
		table:   "user_feedback",
		columns: userFeedbackColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserFeedbackDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserFeedbackDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserFeedbackDao) Columns() UserFeedbackColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserFeedbackDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserFeedbackDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserFeedbackDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
