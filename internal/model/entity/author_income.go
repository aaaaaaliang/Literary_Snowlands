// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthorIncome is the golang structure for table author_income.
type AuthorIncome struct {
	Id             uint64      `json:"id"             ` // 主键
	AuthorId       uint64      `json:"authorId"       ` // 作家ID
	BookId         uint64      `json:"bookId"         ` // 小说ID
	IncomeMonth    *gtime.Time `json:"incomeMonth"    ` // 收入月份
	PreTaxIncome   uint        `json:"preTaxIncome"   ` // 税前收入;单位：分
	AfterTaxIncome uint        `json:"afterTaxIncome" ` // 税后收入;单位：分
	PayStatus      uint        `json:"payStatus"      ` // 支付状态;0-待支付 1-已支付
	ConfirmStatus  uint        `json:"confirmStatus"  ` // 稿费确认状态;0-待确认 1-已确认
	Detail         string      `json:"detail"         ` // 详情
	CreateTime     *gtime.Time `json:"createTime"     ` // 创建时间
	UpdateTime     *gtime.Time `json:"updateTime"     ` // 更新时间
}
