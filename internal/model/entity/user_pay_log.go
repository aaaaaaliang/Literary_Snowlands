// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserPayLog is the golang structure for table user_pay_log.
type UserPayLog struct {
	Id           uint64      `json:"id"           ` //
	UserId       uint64      `json:"userId"       ` // 充值用户ID
	PayChannel   uint        `json:"payChannel"   ` // 充值方式;0-支付宝 1-微信
	OutTradeNo   string      `json:"outTradeNo"   ` // 商户订单号
	Amount       uint        `json:"amount"       ` // 充值金额;单位：分
	ProductType  uint        `json:"productType"  ` // 充值商品类型;0-屋币 1-包年VIP
	ProductId    uint64      `json:"productId"    ` // 充值商品ID
	ProductName  string      `json:"productName"  ` // 充值商品名;示例值：屋币
	ProductValue uint        `json:"productValue" ` // 充值商品值;示例值：255
	PayTime      *gtime.Time `json:"payTime"      ` // 充值时间
	CreateTime   *gtime.Time `json:"createTime"   ` // 创建时间
	UpdateTime   *gtime.Time `json:"updateTime"   ` // 更新时间
}
