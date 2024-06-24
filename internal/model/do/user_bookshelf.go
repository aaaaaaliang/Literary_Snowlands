// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserBookshelf is the golang structure of table user_bookshelf for DAO operations like Where/Data.
type UserBookshelf struct {
	g.Meta       `orm:"table:user_bookshelf, do:true"`
	Id           interface{} // 主键
	UserId       interface{} // 用户ID
	BookId       interface{} // 小说ID
	PreContentId interface{} // 上一次阅读的章节内容表ID
	CreateTime   *gtime.Time // 创建时间;
	UpdateTime   *gtime.Time // 更新时间;
}
