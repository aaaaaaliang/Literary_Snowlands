// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserConsumeLog is the golang structure of table user_consume_log for DAO operations like Where/Data.
type UserConsumeLog struct {
	g.Meta      `orm:"table:user_consume_log, do:true"`
	Id          interface{} // 主键
	UserId      interface{} // 消费用户ID
	Amount      interface{} // 消费使用的金额;单位：屋币
	ProductType interface{} // 消费商品类型;0-小说VIP章节
	ProductId   interface{} // 消费的的商品ID;例如：章节ID
	ProducName  interface{} // 消费的的商品名;例如：章节名
	ProducValue interface{} // 消费的的商品值;例如：1
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 更新时间
}
