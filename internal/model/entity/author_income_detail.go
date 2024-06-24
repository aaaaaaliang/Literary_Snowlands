// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthorIncomeDetail is the golang structure for table author_income_detail.
type AuthorIncomeDetail struct {
	Id            uint64      `json:"id"            ` // 主键
	AuthorId      uint64      `json:"authorId"      ` // 作家ID
	BookId        uint64      `json:"bookId"        ` // 小说ID;0表示全部作品
	IncomeDate    *gtime.Time `json:"incomeDate"    ` // 收入日期
	IncomeAccount uint        `json:"incomeAccount" ` // 订阅总额
	IncomeCount   uint        `json:"incomeCount"   ` // 订阅次数
	IncomeNumber  uint        `json:"incomeNumber"  ` // 订阅人数
	CreateTime    *gtime.Time `json:"createTime"    ` // 创建时间
	UpdateTime    *gtime.Time `json:"updateTime"    ` // 更新时间
}
