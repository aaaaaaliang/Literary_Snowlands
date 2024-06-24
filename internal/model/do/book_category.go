// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BookCategory is the golang structure of table book_category for DAO operations like Where/Data.
type BookCategory struct {
	g.Meta        `orm:"table:book_category, do:true"`
	Id            interface{} //
	WorkDirection interface{} // 作品方向;0-男频 1-女频
	Name          interface{} // 类别名
	Sort          interface{} // 排序
	CreateTime    *gtime.Time // 创建时间
	UpdateTime    *gtime.Time // 更新时间
}
