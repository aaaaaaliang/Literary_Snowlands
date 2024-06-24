// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthorIncomeDetailDao is the data access object for table author_income_detail.
type AuthorIncomeDetailDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns AuthorIncomeDetailColumns // columns contains all the column names of Table for convenient usage.
}

// AuthorIncomeDetailColumns defines and stores column names for table author_income_detail.
type AuthorIncomeDetailColumns struct {
	Id            string // 主键
	AuthorId      string // 作家ID
	BookId        string // 小说ID;0表示全部作品
	IncomeDate    string // 收入日期
	IncomeAccount string // 订阅总额
	IncomeCount   string // 订阅次数
	IncomeNumber  string // 订阅人数
	CreateTime    string // 创建时间
	UpdateTime    string // 更新时间
}

// authorIncomeDetailColumns holds the columns for table author_income_detail.
var authorIncomeDetailColumns = AuthorIncomeDetailColumns{
	Id:            "id",
	AuthorId:      "author_id",
	BookId:        "book_id",
	IncomeDate:    "income_date",
	IncomeAccount: "income_account",
	IncomeCount:   "income_count",
	IncomeNumber:  "income_number",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewAuthorIncomeDetailDao creates and returns a new DAO object for table data access.
func NewAuthorIncomeDetailDao() *AuthorIncomeDetailDao {
	return &AuthorIncomeDetailDao{
		group:   "default",
		table:   "author_income_detail",
		columns: authorIncomeDetailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AuthorIncomeDetailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AuthorIncomeDetailDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AuthorIncomeDetailDao) Columns() AuthorIncomeDetailColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AuthorIncomeDetailDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AuthorIncomeDetailDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AuthorIncomeDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
