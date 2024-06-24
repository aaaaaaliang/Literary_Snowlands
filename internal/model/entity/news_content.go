// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NewsContent is the golang structure for table news_content.
type NewsContent struct {
	Id         uint64      `json:"id"         ` // 主键
	NewsId     uint64      `json:"newsId"     ` // 新闻ID
	Content    string      `json:"content"    ` // 新闻内容
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
}
