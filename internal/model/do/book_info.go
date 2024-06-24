// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BookInfo is the golang structure of table book_info for DAO operations like Where/Data.
type BookInfo struct {
	g.Meta                `orm:"table:book_info, do:true"`
	Id                    interface{} // 主键
	WorkDirection         interface{} // 作品方向;0-男频 1-女频
	CategoryId            interface{} // 类别ID
	CategoryName          interface{} // 类别名
	PicUrl                interface{} // 小说封面地址
	BookName              interface{} // 小说名
	AuthorId              interface{} // 作家id
	AuthorName            interface{} // 作家名
	BookDesc              interface{} // 书籍描述
	Score                 interface{} // 评分;总分:10 ，真实评分 = score/10
	BookStatus            interface{} // 书籍状态;0-连载中 1-已完结
	VisitCount            interface{} // 点击量
	WordCount             interface{} // 总字数
	CommentCount          interface{} // 评论数
	LastChapterId         interface{} // 最新章节ID
	LastChapterName       interface{} // 最新章节名
	LastChapterUpdateTime *gtime.Time // 最新章节更新时间
	IsVip                 interface{} // 是否收费;1-收费 0-免费
	CreateTime            *gtime.Time // 创建时间
	UpdateTime            *gtime.Time // 更新时间
}
