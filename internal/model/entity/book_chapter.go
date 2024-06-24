// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BookChapter is the golang structure for table book_chapter.
type BookChapter struct {
	Id          uint64      `json:"id"          ` //
	BookId      uint64      `json:"bookId"      ` // 小说ID
	ChapterNum  uint        `json:"chapterNum"  ` // 章节号
	ChapterName string      `json:"chapterName" ` // 章节名
	WordCount   uint        `json:"wordCount"   ` // 章节字数
	IsVip       uint        `json:"isVip"       ` // 是否收费;1-收费 0-免费
	CreateTime  *gtime.Time `json:"createTime"  ` //
	UpdateTime  *gtime.Time `json:"updateTime"  ` //
}
