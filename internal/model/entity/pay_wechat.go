// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PayWechat is the golang structure for table pay_wechat.
type PayWechat struct {
	Id             uint64      `json:"id"             ` // 主键
	OutTradeNo     string      `json:"outTradeNo"     ` // 商户订单号
	TransactionId  string      `json:"transactionId"  ` // 微信支付订单号
	TradeType      string      `json:"tradeType"      ` // 交易类型;JSAPI-公众号支付 NATIVE-扫码支付 APP-APP支付 MICROPAY-付款码支付 MWEB-H5支付 FACEPAY-刷脸支付
	TradeState     string      `json:"tradeState"     ` // 交易状态;SUCCESS-支付成功 REFUND-转入退款 NOTPAY-未支付 CLOSED-已关闭 REVOKED-已撤销（付款码支付） USERPAYING-用户支付中（付款码支付） PAYERROR-支付失败(其他原因，如银行返回失败)
	TradeStateDesc string      `json:"tradeStateDesc" ` // 交易状态描述
	Amount         uint        `json:"amount"         ` // 订单总金额;单位：分
	PayerTotal     uint        `json:"payerTotal"     ` // 用户支付金额;单位：分
	SuccessTime    *gtime.Time `json:"successTime"    ` // 支付完成时间
	PayerOpenid    string      `json:"payerOpenid"    ` // 支付者用户标识;用户在直连商户appid下的唯一标识
	CreateTime     *gtime.Time `json:"createTime"     ` // 创建时间
	UpdateTime     *gtime.Time `json:"updateTime"     ` // 更新时间
}
