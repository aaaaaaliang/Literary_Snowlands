// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"Literary_Snowlands/internal/model"
	"Literary_Snowlands/internal/model/entity"
	"context"
)

type (
	IBook interface {
		AddBookVisit(ctx context.Context, bookId int) (err error)
		FindWorkDirectionCategory(ctx context.Context, workDirection int) (category []entity.BookCategory, err error)
		FindBookInfoAndChapterNumById(ctx context.Context, id int) (info *entity.BookInfo, firstChapterID int, err error)
		GetLastChapterAbout(ctx context.Context, id int) (chapterInfo *entity.BookChapter, total int, content string, err error)
		ListRecBooks(ctx context.Context, id int) (re []*entity.BookInfo, err error)
		FindChapterList(ctx context.Context, id int) (chapters []*entity.BookChapter, err error)
		FindBookContentByChapterId(ctx context.Context, id int) (chapterInfo *entity.BookChapter, content string, list *entity.BookInfo, firstChapterId int, err error)
		FindPreChapterId(ctx context.Context, id int) (preId int, err error)
		FindNextChapterId(ctx context.Context, id int) (nextId int, err error)
		ListVisitRank(ctx context.Context) (rank []*entity.BookInfo, err error)
		ListNewestRank(ctx context.Context) (rank []*entity.BookInfo, err error)
		ListUpdateRank(ctx context.Context) (rank []*entity.BookInfo, err error)
		ListNewestComments(ctx context.Context, bookId int) (comments []*model.Comment, commentCount int, err error)
	}
)

var (
	localBook IBook
)

func Book() IBook {
	if localBook == nil {
		panic("implement not found for interface IBook, forgot register?")
	}
	return localBook
}

func RegisterBook(i IBook) {
	localBook = i
}
