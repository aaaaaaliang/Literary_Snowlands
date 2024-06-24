// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package backend

import (
	"context"

	"Literary_Snowlands/api/backend/v1"
)

type IBackendV1 interface {
	GetChapterId(ctx context.Context, req *v1.GetChapterIdReq) (res *v1.GetChapterIdRes, err error)
}

type IAuthorV1 interface {
	AuthorRegister(ctx context.Context, req *v1.AuthorRegisterReq) (res *v1.AuthorRegisterRes, err error)
	QueryAuthorStatus(ctx context.Context, req *v1.QueryAuthorStatusReq) (res *v1.QueryAuthorStatusRes, err error)
	PublishBook(ctx context.Context, req *v1.PublishBookReq) (res *v1.PublishBookRes, err error)
	PublishBookList(ctx context.Context, req *v1.PublishBookListReq) (res *v1.PublishBookListRes, err error)
	PublishBookChapter(ctx context.Context, req *v1.PublishBookChapterReq) (res *v1.PublishBookChapterRes, err error)
	DeleteChapter(ctx context.Context, req *v1.DeleteBookChapterReq) (res *v1.DeleteBookChapterRes, err error)
	GetBookChapter(ctx context.Context, req *v1.GetBookChapterReq) (res *v1.GetBookChapterRes, err error)
	UpdateBookChapter(ctx context.Context, req *v1.UpdateBookChapterReq) (res *v1.UpdateBookChapterRes, err error)
	ListBookChapters(ctx context.Context, req *v1.ListBookChaptersReq) (res *v1.ListBookChaptersRes, err error)
}
