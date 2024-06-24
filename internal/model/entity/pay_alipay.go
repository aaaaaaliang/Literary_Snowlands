// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PayAlipay is the golang structure for table pay_alipay.
type PayAlipay struct {
	Id            uint64      `json:"id"            ` // 主键
	OutTradeNo    string      `json:"outTradeNo"    ` // 商户订单号
	TradeNo       string      `json:"tradeNo"       ` // 支付宝交易号
	BuyerId       string      `json:"buyerId"       ` // 买家支付宝账号 ID
	TradeStatus   string      `json:"tradeStatus"   ` // 交易状态;TRADE_SUCCESS-交易成功
	TotalAmount   uint        `json:"totalAmount"   ` // 订单金额;单位：分
	ReceiptAmount uint        `json:"receiptAmount" ` // 实收金额;单位：分
	InvoiceAmount uint        `json:"invoiceAmount" ` // 开票金额
	GmtCreate     *gtime.Time `json:"gmtCreate"     ` // 交易创建时间
	GmtPayment    *gtime.Time `json:"gmtPayment"    ` // 交易付款时间
	CreateTime    *gtime.Time `json:"createTime"    ` // 创建时间
	UpdateTime    *gtime.Time `json:"updateTime"    ` // 更新时间
}
