// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserPayLog is the golang structure of table user_pay_log for DAO operations like Where/Data.
type UserPayLog struct {
	g.Meta       `orm:"table:user_pay_log, do:true"`
	Id           interface{} //
	UserId       interface{} // 充值用户ID
	PayChannel   interface{} // 充值方式;0-支付宝 1-微信
	OutTradeNo   interface{} // 商户订单号
	Amount       interface{} // 充值金额;单位：分
	ProductType  interface{} // 充值商品类型;0-屋币 1-包年VIP
	ProductId    interface{} // 充值商品ID
	ProductName  interface{} // 充值商品名;示例值：屋币
	ProductValue interface{} // 充值商品值;示例值：255
	PayTime      *gtime.Time // 充值时间
	CreateTime   *gtime.Time // 创建时间
	UpdateTime   *gtime.Time // 更新时间
}
