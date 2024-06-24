// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserConsumeLog is the golang structure for table user_consume_log.
type UserConsumeLog struct {
	Id          uint64      `json:"id"          ` // 主键
	UserId      uint64      `json:"userId"      ` // 消费用户ID
	Amount      uint        `json:"amount"      ` // 消费使用的金额;单位：屋币
	ProductType uint        `json:"productType" ` // 消费商品类型;0-小说VIP章节
	ProductId   uint64      `json:"productId"   ` // 消费的的商品ID;例如：章节ID
	ProducName  string      `json:"producName"  ` // 消费的的商品名;例如：章节名
	ProducValue uint        `json:"producValue" ` // 消费的的商品值;例如：1
	CreateTime  *gtime.Time `json:"createTime"  ` // 创建时间
	UpdateTime  *gtime.Time `json:"updateTime"  ` // 更新时间
}
