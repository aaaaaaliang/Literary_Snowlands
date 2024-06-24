// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthorInfoDao is the data access object for table author_info.
type AuthorInfoDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns AuthorInfoColumns // columns contains all the column names of Table for convenient usage.
}

// AuthorInfoColumns defines and stores column names for table author_info.
type AuthorInfoColumns struct {
	Id            string // 主键
	UserId        string // 用户ID
	InviteCode    string // 邀请码
	PenName       string // 笔名
	TelPhone      string // 手机号码
	ChatAccount   string // QQ或微信账号
	Email         string // 电子邮箱
	WorkDirection string // 作品方向;0-男频 1-女频
	Status        string // 0：正常;1-封禁
	CreateTime    string // 创建时间
	UpdateTime    string // 更新时间
}

// authorInfoColumns holds the columns for table author_info.
var authorInfoColumns = AuthorInfoColumns{
	Id:            "id",
	UserId:        "user_id",
	InviteCode:    "invite_code",
	PenName:       "pen_name",
	TelPhone:      "tel_phone",
	ChatAccount:   "chat_account",
	Email:         "email",
	WorkDirection: "work_direction",
	Status:        "status",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewAuthorInfoDao creates and returns a new DAO object for table data access.
func NewAuthorInfoDao() *AuthorInfoDao {
	return &AuthorInfoDao{
		group:   "default",
		table:   "author_info",
		columns: authorInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AuthorInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AuthorInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AuthorInfoDao) Columns() AuthorInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AuthorInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AuthorInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AuthorInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
