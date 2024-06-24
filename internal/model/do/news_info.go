// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NewsInfo is the golang structure of table news_info for DAO operations like Where/Data.
type NewsInfo struct {
	g.Meta       `orm:"table:news_info, do:true"`
	Id           interface{} // 主键
	CategoryId   interface{} // 类别ID
	CategoryName interface{} // 类别名
	SourceName   interface{} // 新闻来源
	Title        interface{} // 新闻标题
	CreateTime   *gtime.Time // 创建时间
	UpdateTime   *gtime.Time // 更新时间
}
