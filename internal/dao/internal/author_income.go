// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthorIncomeDao is the data access object for table author_income.
type AuthorIncomeDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns AuthorIncomeColumns // columns contains all the column names of Table for convenient usage.
}

// AuthorIncomeColumns defines and stores column names for table author_income.
type AuthorIncomeColumns struct {
	Id             string // 主键
	AuthorId       string // 作家ID
	BookId         string // 小说ID
	IncomeMonth    string // 收入月份
	PreTaxIncome   string // 税前收入;单位：分
	AfterTaxIncome string // 税后收入;单位：分
	PayStatus      string // 支付状态;0-待支付 1-已支付
	ConfirmStatus  string // 稿费确认状态;0-待确认 1-已确认
	Detail         string // 详情
	CreateTime     string // 创建时间
	UpdateTime     string // 更新时间
}

// authorIncomeColumns holds the columns for table author_income.
var authorIncomeColumns = AuthorIncomeColumns{
	Id:             "id",
	AuthorId:       "author_id",
	BookId:         "book_id",
	IncomeMonth:    "income_month",
	PreTaxIncome:   "pre_tax_income",
	AfterTaxIncome: "after_tax_income",
	PayStatus:      "pay_status",
	ConfirmStatus:  "confirm_status",
	Detail:         "detail",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
}

// NewAuthorIncomeDao creates and returns a new DAO object for table data access.
func NewAuthorIncomeDao() *AuthorIncomeDao {
	return &AuthorIncomeDao{
		group:   "default",
		table:   "author_income",
		columns: authorIncomeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AuthorIncomeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AuthorIncomeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AuthorIncomeDao) Columns() AuthorIncomeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AuthorIncomeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AuthorIncomeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AuthorIncomeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
