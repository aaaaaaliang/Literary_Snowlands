// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PayAlipayDao is the data access object for table pay_alipay.
type PayAlipayDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns PayAlipayColumns // columns contains all the column names of Table for convenient usage.
}

// PayAlipayColumns defines and stores column names for table pay_alipay.
type PayAlipayColumns struct {
	Id            string // 主键
	OutTradeNo    string // 商户订单号
	TradeNo       string // 支付宝交易号
	BuyerId       string // 买家支付宝账号 ID
	TradeStatus   string // 交易状态;TRADE_SUCCESS-交易成功
	TotalAmount   string // 订单金额;单位：分
	ReceiptAmount string // 实收金额;单位：分
	InvoiceAmount string // 开票金额
	GmtCreate     string // 交易创建时间
	GmtPayment    string // 交易付款时间
	CreateTime    string // 创建时间
	UpdateTime    string // 更新时间
}

// payAlipayColumns holds the columns for table pay_alipay.
var payAlipayColumns = PayAlipayColumns{
	Id:            "id",
	OutTradeNo:    "out_trade_no",
	TradeNo:       "trade_no",
	BuyerId:       "buyer_id",
	TradeStatus:   "trade_status",
	TotalAmount:   "total_amount",
	ReceiptAmount: "receipt_amount",
	InvoiceAmount: "invoice_amount",
	GmtCreate:     "gmt_create",
	GmtPayment:    "gmt_payment",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewPayAlipayDao creates and returns a new DAO object for table data access.
func NewPayAlipayDao() *PayAlipayDao {
	return &PayAlipayDao{
		group:   "default",
		table:   "pay_alipay",
		columns: payAlipayColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PayAlipayDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PayAlipayDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PayAlipayDao) Columns() PayAlipayColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PayAlipayDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PayAlipayDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PayAlipayDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
