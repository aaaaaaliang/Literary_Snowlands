// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HomeBook is the golang structure of table home_book for DAO operations like Where/Data.
type HomeBook struct {
	g.Meta     `orm:"table:home_book, do:true"`
	Id         interface{} //
	Type       interface{} // 推荐类型;0-轮播图 1-顶部栏 2-本周强推 3-热门推荐 4-精品推荐
	Sort       interface{} // 推荐排序
	BookId     interface{} // 推荐小说ID
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
}
