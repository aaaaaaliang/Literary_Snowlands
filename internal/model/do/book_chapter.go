// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BookChapter is the golang structure of table book_chapter for DAO operations like Where/Data.
type BookChapter struct {
	g.Meta      `orm:"table:book_chapter, do:true"`
	Id          interface{} //
	BookId      interface{} // 小说ID
	ChapterNum  interface{} // 章节号
	ChapterName interface{} // 章节名
	WordCount   interface{} // 章节字数
	IsVip       interface{} // 是否收费;1-收费 0-免费
	CreateTime  *gtime.Time //
	UpdateTime  *gtime.Time //
}
