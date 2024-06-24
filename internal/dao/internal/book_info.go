// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BookInfoDao is the data access object for table book_info.
type BookInfoDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns BookInfoColumns // columns contains all the column names of Table for convenient usage.
}

// BookInfoColumns defines and stores column names for table book_info.
type BookInfoColumns struct {
	Id                    string // 主键
	WorkDirection         string // 作品方向;0-男频 1-女频
	CategoryId            string // 类别ID
	CategoryName          string // 类别名
	PicUrl                string // 小说封面地址
	BookName              string // 小说名
	AuthorId              string // 作家id
	AuthorName            string // 作家名
	BookDesc              string // 书籍描述
	Score                 string // 评分;总分:10 ，真实评分 = score/10
	BookStatus            string // 书籍状态;0-连载中 1-已完结
	VisitCount            string // 点击量
	WordCount             string // 总字数
	CommentCount          string // 评论数
	LastChapterId         string // 最新章节ID
	LastChapterName       string // 最新章节名
	LastChapterUpdateTime string // 最新章节更新时间
	IsVip                 string // 是否收费;1-收费 0-免费
	CreateTime            string // 创建时间
	UpdateTime            string // 更新时间
}

// bookInfoColumns holds the columns for table book_info.
var bookInfoColumns = BookInfoColumns{
	Id:                    "id",
	WorkDirection:         "work_direction",
	CategoryId:            "category_id",
	CategoryName:          "category_name",
	PicUrl:                "pic_url",
	BookName:              "book_name",
	AuthorId:              "author_id",
	AuthorName:            "author_name",
	BookDesc:              "book_desc",
	Score:                 "score",
	BookStatus:            "book_status",
	VisitCount:            "visit_count",
	WordCount:             "word_count",
	CommentCount:          "comment_count",
	LastChapterId:         "last_chapter_id",
	LastChapterName:       "last_chapter_name",
	LastChapterUpdateTime: "last_chapter_update_time",
	IsVip:                 "is_vip",
	CreateTime:            "create_time",
	UpdateTime:            "update_time",
}

// NewBookInfoDao creates and returns a new DAO object for table data access.
func NewBookInfoDao() *BookInfoDao {
	return &BookInfoDao{
		group:   "default",
		table:   "book_info",
		columns: bookInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BookInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BookInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BookInfoDao) Columns() BookInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BookInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BookInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BookInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
