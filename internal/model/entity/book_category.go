// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BookCategory is the golang structure for table book_category.
type BookCategory struct {
	Id            uint64      `json:"id"            ` //
	WorkDirection uint        `json:"workDirection" ` // 作品方向;0-男频 1-女频
	Name          string      `json:"name"          ` // 类别名
	Sort          uint        `json:"sort"          ` // 排序
	CreateTime    *gtime.Time `json:"createTime"    ` // 创建时间
	UpdateTime    *gtime.Time `json:"updateTime"    ` // 更新时间
}
