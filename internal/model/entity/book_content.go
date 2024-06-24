// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BookContent is the golang structure for table book_content.
type BookContent struct {
	Id         uint64      `json:"id"         ` // 主键
	ChapterId  uint64      `json:"chapterId"  ` // 章节ID
	Content    string      `json:"content"    ` // 小说章节内容
	CreateTime *gtime.Time `json:"createTime" ` //
	UpdateTime *gtime.Time `json:"updateTime" ` //
}
