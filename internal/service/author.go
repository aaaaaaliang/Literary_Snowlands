// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"Literary_Snowlands/internal/model"
	"Literary_Snowlands/internal/model/do"
	"Literary_Snowlands/internal/model/entity"
	"context"
)

type IAuthor interface {
	RegisterAuthor(ctx context.Context, registerInfo do.AuthorInfo) error
	QueryAuthorStatus(ctx context.Context) (*int, error)
	CreateBook(ctx context.Context, bookInfo do.BookInfo) error
	ListAuthorBookInfo(ctx context.Context, pageNum, pageSize int) (bookInfoList []entity.BookInfo, total int, err error)
	SaveBookChapter(ctx context.Context, input model.AuthorBookInput) error
	DeleteChapterById(ctx context.Context, chapterId string) error
	GetChapterById(ctx context.Context, chapterId string) (*model.BookChapter, error)
	UpdateChapterById(ctx context.Context, chapterId string, chapter model.BookChapter) error
	ListBookChapters(ctx context.Context, book model.ListBook) (listBookChapters []entity.BookChapter, total int, err error)
}

var (
	localAuthor IAuthor
)

func Author() IAuthor {
	if localAuthor == nil {
		panic("implement not found for interface IAuthor, forgot register?")
	}
	return localAuthor
}

func RegisterAuthor(i IAuthor) {
	localAuthor = i
}
