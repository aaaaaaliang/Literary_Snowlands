// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NewsInfo is the golang structure for table news_info.
type NewsInfo struct {
	Id           uint64      `json:"id"           ` // 主键
	CategoryId   uint64      `json:"categoryId"   ` // 类别ID
	CategoryName string      `json:"categoryName" ` // 类别名
	SourceName   string      `json:"sourceName"   ` // 新闻来源
	Title        string      `json:"title"        ` // 新闻标题
	CreateTime   *gtime.Time `json:"createTime"   ` // 创建时间
	UpdateTime   *gtime.Time `json:"updateTime"   ` // 更新时间
}
