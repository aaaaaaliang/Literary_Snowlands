// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NewsCategory is the golang structure of table news_category for DAO operations like Where/Data.
type NewsCategory struct {
	g.Meta     `orm:"table:news_category, do:true"`
	Id         interface{} //
	Name       interface{} // 类别名
	Sort       interface{} // 排序
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
}
