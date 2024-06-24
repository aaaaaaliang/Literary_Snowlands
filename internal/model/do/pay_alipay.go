// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PayAlipay is the golang structure of table pay_alipay for DAO operations like Where/Data.
type PayAlipay struct {
	g.Meta        `orm:"table:pay_alipay, do:true"`
	Id            interface{} // 主键
	OutTradeNo    interface{} // 商户订单号
	TradeNo       interface{} // 支付宝交易号
	BuyerId       interface{} // 买家支付宝账号 ID
	TradeStatus   interface{} // 交易状态;TRADE_SUCCESS-交易成功
	TotalAmount   interface{} // 订单金额;单位：分
	ReceiptAmount interface{} // 实收金额;单位：分
	InvoiceAmount interface{} // 开票金额
	GmtCreate     *gtime.Time // 交易创建时间
	GmtPayment    *gtime.Time // 交易付款时间
	CreateTime    *gtime.Time // 创建时间
	UpdateTime    *gtime.Time // 更新时间
}
