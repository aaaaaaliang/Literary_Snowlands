// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package frontend

import (
	"context"

	"Literary_Snowlands/api/frontend/v1"
)

type IFrontendV1 interface {
	// 小说模块
	AddBookVisit(ctx context.Context, req *v1.AddBookVisitReq) (res *v1.AddBookVisitRes, err error)
	ListBookCategory(ctx context.Context, req *v1.ListBookCategoryReq) (res []*v1.ListBookCategoryRes, err error)
	GetBookInfoById(ctx context.Context, req *v1.GetBookInfoByIdReq) (res *v1.GetBookInfoByIdRes, err error)
	GetLastChapterAbout(ctx context.Context, req *v1.GetLastChapterAboutReq) (res *v1.GetLastChapterAboutRes, err error)
	ListRecommendedBooks(ctx context.Context, req *v1.ListRecommendedBooksReq) (res []*v1.ListRecommendedBooksRes, err error)
	ListChapters(ctx context.Context, req *v1.ListChaptersReq) (res []*v1.ListChaptersRes, err error)
	GetBookContentAbout(ctx context.Context, req *v1.GetBookContentAboutReq) (res *v1.GetBookContentAboutRes, err error)
	GetPreChapterId(ctx context.Context, req *v1.GetPreChapterIdReq) (res *v1.GetPreChapterIdRes, err error)
	GetNextChapterId(ctx context.Context, req *v1.GetNextChapterIdReq) (res *v1.GetNextChapterIdRes, err error)
	GetBookVisitRank(ctx context.Context, req *v1.GetBookVisitRankReq) (res []*v1.GetBookVisitRankRes, err error)
	GetBookNewestRank(ctx context.Context, req *v1.GetBookNewestRankReq) (res []*v1.GetBookNewestRankRes, err error)
	GetBookUpdateRank(ctx context.Context, req *v1.GetBookUpdateRankReq) (res []*v1.GetBookUpdateRankRes, err error)
	ListBookComment(ctx context.Context, req *v1.ListBookCommentReq) (res *v1.ListBookCommentRes, err error)
}

type IUserV1 interface {
	// 用户模块
	GetCaptcha(ctx context.Context, req *v1.GetCaptchaReq) (res *v1.GetCaptchaRes, err error)
	SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error)
	SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error)
}

type IHomeV1 interface {
	QueryRecommendedBooks(ctx context.Context, req *v1.QueryRecommendedBooksReq) (res []*v1.QueryRecommendedBooksRes, err error)
	QueryFriendLinkBooks(ctx context.Context, req *v1.QueryFriendLinkBooksReq) (res []*v1.QueryFriendLinkBooksRes, err error)
}

type INewsV1 interface {
	QueryNewsInfo(ctx context.Context, req *v1.QueryNewsInfoReq) (res *v1.QueryNewsInfoRes, err error)
	QueryNewsLatestInfo(ctx context.Context, req *v1.QueryNewsLatestInfoReq) (res []*v1.QueryNewsLatestInfoRes, err error)
}

type ISearchV1 interface {
	SearchBooks(ctx context.Context, req *v1.SearchBooksReq) (res *v1.SearchBooksRes, err error)
}

type IUtilV1 interface {
	ResourceImages(ctx context.Context, req *v1.ResourceImageReq) (res *v1.ResourceImageRes, err error)
}
