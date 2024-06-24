package frontend

import (
	"Literary_Snowlands/api/frontend"
	v1 "Literary_Snowlands/api/frontend/v1"
	"Literary_Snowlands/internal/service"
	"context"
)

type HomeControllerV1 struct {
}

func NewHomeControllerV1() frontend.IHomeV1 {
	return &HomeControllerV1{}
}

func (c *HomeControllerV1) QueryFriendLinkBooks(ctx context.Context, req *v1.QueryFriendLinkBooksReq) (res []*v1.QueryFriendLinkBooksRes, err error) {
	friendLinkList, err := service.Home().QueryFriendLinkList(ctx)
	if err != nil {
		return nil, err
	}
	var friendLinkListArrayRes []*v1.QueryFriendLinkBooksRes
	for _, v := range friendLinkList {
		friendLinkListArrayRes = append(friendLinkListArrayRes, &v1.QueryFriendLinkBooksRes{
			LinkName: v.LinkName,
			LinkUrl:  v.LinkUrl,
		})
	}

	return friendLinkListArrayRes, nil
}

func (c *HomeControllerV1) QueryRecommendedBooks(ctx context.Context, req *v1.QueryRecommendedBooksReq) (res []*v1.QueryRecommendedBooksRes, err error) {
	booksListArray, err := service.Home().FindRecommendBook(ctx)
	if err != nil {
		return nil, err
	}
	var booksListResArray []*v1.QueryRecommendedBooksRes
	for _, v := range booksListArray {
		booksListResArray = append(booksListResArray, &v1.QueryRecommendedBooksRes{
			Type:       v.Type,
			BookId:     v.BookId,
			PicUrl:     v.PicUrl,
			BookName:   v.BookName,
			AuthorName: v.AuthorName,
			BookDesc:   v.BookDesc,
		})
	}

	return booksListResArray, nil
}
