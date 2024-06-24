// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NewsCategory is the golang structure for table news_category.
type NewsCategory struct {
	Id         uint64      `json:"id"         ` //
	Name       string      `json:"name"       ` // 类别名
	Sort       uint        `json:"sort"       ` // 排序
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
}
