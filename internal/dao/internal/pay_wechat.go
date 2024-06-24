// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PayWechatDao is the data access object for table pay_wechat.
type PayWechatDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns PayWechatColumns // columns contains all the column names of Table for convenient usage.
}

// PayWechatColumns defines and stores column names for table pay_wechat.
type PayWechatColumns struct {
	Id             string // 主键
	OutTradeNo     string // 商户订单号
	TransactionId  string // 微信支付订单号
	TradeType      string // 交易类型;JSAPI-公众号支付 NATIVE-扫码支付 APP-APP支付 MICROPAY-付款码支付 MWEB-H5支付 FACEPAY-刷脸支付
	TradeState     string // 交易状态;SUCCESS-支付成功 REFUND-转入退款 NOTPAY-未支付 CLOSED-已关闭 REVOKED-已撤销（付款码支付） USERPAYING-用户支付中（付款码支付） PAYERROR-支付失败(其他原因，如银行返回失败)
	TradeStateDesc string // 交易状态描述
	Amount         string // 订单总金额;单位：分
	PayerTotal     string // 用户支付金额;单位：分
	SuccessTime    string // 支付完成时间
	PayerOpenid    string // 支付者用户标识;用户在直连商户appid下的唯一标识
	CreateTime     string // 创建时间
	UpdateTime     string // 更新时间
}

// payWechatColumns holds the columns for table pay_wechat.
var payWechatColumns = PayWechatColumns{
	Id:             "id",
	OutTradeNo:     "out_trade_no",
	TransactionId:  "transaction_id",
	TradeType:      "trade_type",
	TradeState:     "trade_state",
	TradeStateDesc: "trade_state_desc",
	Amount:         "amount",
	PayerTotal:     "payer_total",
	SuccessTime:    "success_time",
	PayerOpenid:    "payer_openid",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
}

// NewPayWechatDao creates and returns a new DAO object for table data access.
func NewPayWechatDao() *PayWechatDao {
	return &PayWechatDao{
		group:   "default",
		table:   "pay_wechat",
		columns: payWechatColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PayWechatDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PayWechatDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PayWechatDao) Columns() PayWechatColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PayWechatDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PayWechatDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PayWechatDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
