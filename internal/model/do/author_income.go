// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthorIncome is the golang structure of table author_income for DAO operations like Where/Data.
type AuthorIncome struct {
	g.Meta         `orm:"table:author_income, do:true"`
	Id             interface{} // 主键
	AuthorId       interface{} // 作家ID
	BookId         interface{} // 小说ID
	IncomeMonth    *gtime.Time // 收入月份
	PreTaxIncome   interface{} // 税前收入;单位：分
	AfterTaxIncome interface{} // 税后收入;单位：分
	PayStatus      interface{} // 支付状态;0-待支付 1-已支付
	ConfirmStatus  interface{} // 稿费确认状态;0-待确认 1-已确认
	Detail         interface{} // 详情
	CreateTime     *gtime.Time // 创建时间
	UpdateTime     *gtime.Time // 更新时间
}
