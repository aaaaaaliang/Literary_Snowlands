// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthorIncomeDetail is the golang structure of table author_income_detail for DAO operations like Where/Data.
type AuthorIncomeDetail struct {
	g.Meta        `orm:"table:author_income_detail, do:true"`
	Id            interface{} // 主键
	AuthorId      interface{} // 作家ID
	BookId        interface{} // 小说ID;0表示全部作品
	IncomeDate    *gtime.Time // 收入日期
	IncomeAccount interface{} // 订阅总额
	IncomeCount   interface{} // 订阅次数
	IncomeNumber  interface{} // 订阅人数
	CreateTime    *gtime.Time // 创建时间
	UpdateTime    *gtime.Time // 更新时间
}
