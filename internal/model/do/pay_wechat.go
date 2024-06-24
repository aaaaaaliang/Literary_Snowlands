// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PayWechat is the golang structure of table pay_wechat for DAO operations like Where/Data.
type PayWechat struct {
	g.Meta         `orm:"table:pay_wechat, do:true"`
	Id             interface{} // 主键
	OutTradeNo     interface{} // 商户订单号
	TransactionId  interface{} // 微信支付订单号
	TradeType      interface{} // 交易类型;JSAPI-公众号支付 NATIVE-扫码支付 APP-APP支付 MICROPAY-付款码支付 MWEB-H5支付 FACEPAY-刷脸支付
	TradeState     interface{} // 交易状态;SUCCESS-支付成功 REFUND-转入退款 NOTPAY-未支付 CLOSED-已关闭 REVOKED-已撤销（付款码支付） USERPAYING-用户支付中（付款码支付） PAYERROR-支付失败(其他原因，如银行返回失败)
	TradeStateDesc interface{} // 交易状态描述
	Amount         interface{} // 订单总金额;单位：分
	PayerTotal     interface{} // 用户支付金额;单位：分
	SuccessTime    *gtime.Time // 支付完成时间
	PayerOpenid    interface{} // 支付者用户标识;用户在直连商户appid下的唯一标识
	CreateTime     *gtime.Time // 创建时间
	UpdateTime     *gtime.Time // 更新时间
}
