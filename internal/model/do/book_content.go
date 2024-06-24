// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BookContent is the golang structure of table book_content for DAO operations like Where/Data.
type BookContent struct {
	g.Meta     `orm:"table:book_content, do:true"`
	Id         interface{} // 主键
	ChapterId  interface{} // 章节ID
	Content    interface{} // 小说章节内容
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
}
