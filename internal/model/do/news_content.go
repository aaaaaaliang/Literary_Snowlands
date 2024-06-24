// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NewsContent is the golang structure of table news_content for DAO operations like Where/Data.
type NewsContent struct {
	g.Meta     `orm:"table:news_content, do:true"`
	Id         interface{} // 主键
	NewsId     interface{} // 新闻ID
	Content    interface{} // 新闻内容
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
}
