// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BookChapterDao is the data access object for table book_chapter.
type BookChapterDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns BookChapterColumns // columns contains all the column names of Table for convenient usage.
}

// BookChapterColumns defines and stores column names for table book_chapter.
type BookChapterColumns struct {
	Id          string //
	BookId      string // 小说ID
	ChapterNum  string // 章节号
	ChapterName string // 章节名
	WordCount   string // 章节字数
	IsVip       string // 是否收费;1-收费 0-免费
	CreateTime  string //
	UpdateTime  string //
}

// bookChapterColumns holds the columns for table book_chapter.
var bookChapterColumns = BookChapterColumns{
	Id:          "id",
	BookId:      "book_id",
	ChapterNum:  "chapter_num",
	ChapterName: "chapter_name",
	WordCount:   "word_count",
	IsVip:       "is_vip",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

// NewBookChapterDao creates and returns a new DAO object for table data access.
func NewBookChapterDao() *BookChapterDao {
	return &BookChapterDao{
		group:   "default",
		table:   "book_chapter",
		columns: bookChapterColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BookChapterDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BookChapterDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BookChapterDao) Columns() BookChapterColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BookChapterDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BookChapterDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BookChapterDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
