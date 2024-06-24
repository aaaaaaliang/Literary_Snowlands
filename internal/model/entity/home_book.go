// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HomeBook is the golang structure for table home_book.
type HomeBook struct {
	Id         uint64      `json:"id"         ` //
	Type       uint        `json:"type"       ` // 推荐类型;0-轮播图 1-顶部栏 2-本周强推 3-热门推荐 4-精品推荐
	Sort       uint        `json:"sort"       ` // 推荐排序
	BookId     uint64      `json:"bookId"     ` // 推荐小说ID
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
}
