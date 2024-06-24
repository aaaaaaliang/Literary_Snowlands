// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BookInfo is the golang structure for table book_info.
type BookInfo struct {
	Id                    uint64      `json:"id"                    ` // 主键
	WorkDirection         uint        `json:"workDirection"         ` // 作品方向;0-男频 1-女频
	CategoryId            uint64      `json:"categoryId"            ` // 类别ID
	CategoryName          string      `json:"categoryName"          ` // 类别名
	PicUrl                string      `json:"picUrl"                ` // 小说封面地址
	BookName              string      `json:"bookName"              ` // 小说名
	AuthorId              uint64      `json:"authorId"              ` // 作家id
	AuthorName            string      `json:"authorName"            ` // 作家名
	BookDesc              string      `json:"bookDesc"              ` // 书籍描述
	Score                 uint        `json:"score"                 ` // 评分;总分:10 ，真实评分 = score/10
	BookStatus            uint        `json:"bookStatus"            ` // 书籍状态;0-连载中 1-已完结
	VisitCount            uint64      `json:"visitCount"            ` // 点击量
	WordCount             uint        `json:"wordCount"             ` // 总字数
	CommentCount          uint        `json:"commentCount"          ` // 评论数
	LastChapterId         uint64      `json:"lastChapterId"         ` // 最新章节ID
	LastChapterName       string      `json:"lastChapterName"       ` // 最新章节名
	LastChapterUpdateTime *gtime.Time `json:"lastChapterUpdateTime" ` // 最新章节更新时间
	IsVip                 uint        `json:"isVip"                 ` // 是否收费;1-收费 0-免费
	CreateTime            *gtime.Time `json:"createTime"            ` // 创建时间
	UpdateTime            *gtime.Time `json:"updateTime"            ` // 更新时间
}
