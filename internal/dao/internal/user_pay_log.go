// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserPayLogDao is the data access object for table user_pay_log.
type UserPayLogDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns UserPayLogColumns // columns contains all the column names of Table for convenient usage.
}

// UserPayLogColumns defines and stores column names for table user_pay_log.
type UserPayLogColumns struct {
	Id           string //
	UserId       string // 充值用户ID
	PayChannel   string // 充值方式;0-支付宝 1-微信
	OutTradeNo   string // 商户订单号
	Amount       string // 充值金额;单位：分
	ProductType  string // 充值商品类型;0-屋币 1-包年VIP
	ProductId    string // 充值商品ID
	ProductName  string // 充值商品名;示例值：屋币
	ProductValue string // 充值商品值;示例值：255
	PayTime      string // 充值时间
	CreateTime   string // 创建时间
	UpdateTime   string // 更新时间
}

// userPayLogColumns holds the columns for table user_pay_log.
var userPayLogColumns = UserPayLogColumns{
	Id:           "id",
	UserId:       "user_id",
	PayChannel:   "pay_channel",
	OutTradeNo:   "out_trade_no",
	Amount:       "amount",
	ProductType:  "product_type",
	ProductId:    "product_id",
	ProductName:  "product_name",
	ProductValue: "product_value",
	PayTime:      "pay_time",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
}

// NewUserPayLogDao creates and returns a new DAO object for table data access.
func NewUserPayLogDao() *UserPayLogDao {
	return &UserPayLogDao{
		group:   "default",
		table:   "user_pay_log",
		columns: userPayLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserPayLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserPayLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserPayLogDao) Columns() UserPayLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserPayLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserPayLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserPayLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
