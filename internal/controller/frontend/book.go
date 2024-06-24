package frontend

import (
	"Literary_Snowlands/api/frontend"
	v1 "Literary_Snowlands/api/frontend/v1"
	"Literary_Snowlands/internal/service"
	"context"
	"strconv"
)

type BookControllerV1 struct{}

func NewBookControllerV1() frontend.IFrontendV1 {
	return &BookControllerV1{}
}

func (c *BookControllerV1) ListChapters(ctx context.Context, req *v1.ListChaptersReq) (res []*v1.ListChaptersRes, err error) {

	chapterListInfo, err := service.Book().FindChapterList(ctx, req.BookId)
	if err != nil {
		return nil, err
	}
	var chapterListInfoRes []*v1.ListChaptersRes
	for _, v := range chapterListInfo {
		chapterListInfoRes = append(chapterListInfoRes, &v1.ListChaptersRes{
			Id:                strconv.FormatUint(v.Id, 10),
			BookId:            strconv.FormatUint(v.BookId, 10),
			ChapterNum:        v.ChapterNum,
			ChapterName:       v.ChapterName,
			ChapterWordCount:  v.WordCount,
			ChapterUpdateTime: v.UpdateTime,
			IsVip:             v.IsVip,
		})
	}

	return chapterListInfoRes, nil
}

func (c *BookControllerV1) GetBookContentAbout(ctx context.Context, req *v1.GetBookContentAboutReq) (res *v1.GetBookContentAboutRes, err error) {

	chapterInfo, content, bookInfo, firstChapterId, err := service.Book().FindBookContentByChapterId(ctx, req.ChapterId)
	if err != nil {
		return nil, err
	}
	bookInfoV1 := &v1.BookInfo{
		Id:              strconv.FormatUint(bookInfo.Id, 10),
		CategoryId:      bookInfo.CategoryId,
		CategoryName:    bookInfo.CategoryName,
		PicUrl:          bookInfo.PicUrl,
		BookName:        bookInfo.BookName,
		AuthorId:        bookInfo.AuthorId,
		AuthorName:      bookInfo.AuthorName,
		BookDesc:        bookInfo.BookDesc,
		BookStatus:      bookInfo.BookStatus,
		VisitCount:      bookInfo.VisitCount,
		WordCount:       bookInfo.WordCount,
		CommentCount:    bookInfo.CommentCount,
		FirstChapterId:  strconv.Itoa(firstChapterId),
		LastChapterId:   strconv.FormatUint(bookInfo.LastChapterId, 10),
		LastChapterName: bookInfo.LastChapterName,
		UpdateTime:      bookInfo.UpdateTime,
	}

	chapterInfoV1 := &v1.ChapterInfo{
		Id:                strconv.FormatUint(chapterInfo.Id, 10),
		BookId:            strconv.FormatUint(chapterInfo.BookId, 10),
		ChapterNum:        chapterInfo.ChapterNum,
		ChapterName:       chapterInfo.ChapterName,
		ChapterWordCount:  chapterInfo.WordCount,
		ChapterUpdateTime: chapterInfo.UpdateTime,
		IsVip:             chapterInfo.IsVip,
	}

	return &v1.GetBookContentAboutRes{
		BookInfo:    *bookInfoV1,
		ChapterInfo: *chapterInfoV1,
		BookContent: content,
	}, nil
}

func (c *BookControllerV1) GetBookInfoById(ctx context.Context, req *v1.GetBookInfoByIdReq) (res *v1.GetBookInfoByIdRes, err error) {
	bookInfo, firstChapterId, err := service.Book().FindBookInfoAndChapterNumById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	var bookInfoRes v1.BookInfo
	bookInfoRes.Id = strconv.FormatUint(bookInfo.Id, 10)
	bookInfoRes.CategoryId = bookInfo.CategoryId
	bookInfoRes.CategoryName = bookInfo.CategoryName
	bookInfoRes.PicUrl = bookInfo.PicUrl
	bookInfoRes.BookName = bookInfo.BookName
	bookInfoRes.AuthorId = bookInfo.AuthorId
	bookInfoRes.AuthorName = bookInfo.AuthorName
	bookInfoRes.BookDesc = bookInfo.BookDesc
	bookInfoRes.BookStatus = bookInfo.BookStatus
	bookInfoRes.VisitCount = bookInfo.VisitCount
	bookInfoRes.WordCount = bookInfo.WordCount
	bookInfoRes.CommentCount = bookInfo.CommentCount
	bookInfoRes.FirstChapterId = strconv.Itoa(firstChapterId)
	bookInfoRes.LastChapterId = strconv.FormatUint(bookInfo.LastChapterId, 10)
	bookInfoRes.LastChapterName = bookInfo.LastChapterName
	bookInfoRes.UpdateTime = bookInfo.UpdateTime

	return &v1.GetBookInfoByIdRes{BookInfo: bookInfoRes}, nil
}

func (c *BookControllerV1) GetLastChapterAbout(ctx context.Context, req *v1.GetLastChapterAboutReq) (res *v1.GetLastChapterAboutRes, err error) {
	chapterInfo, total, content, err := service.Book().GetLastChapterAbout(ctx, req.BookId)
	if err != nil {
		return nil, err
	}
	if chapterInfo == nil {
		return nil, nil
	}
	chapterReq := v1.ChapterInfo{
		Id:                strconv.FormatUint(chapterInfo.Id, 10),
		BookId:            strconv.FormatUint(chapterInfo.BookId, 10),
		ChapterNum:        chapterInfo.ChapterNum,
		ChapterName:       chapterInfo.ChapterName,
		ChapterWordCount:  chapterInfo.WordCount,
		ChapterUpdateTime: chapterInfo.UpdateTime,
		IsVip:             chapterInfo.IsVip,
	}

	LastChapterInfo := v1.GetLastChapterAboutRes{
		ChapterInfo:    chapterReq,
		ChapterTotal:   total,
		ContentSummary: content,
	}
	return &LastChapterInfo, nil
}

func (c *BookControllerV1) ListBookComment(ctx context.Context, req *v1.ListBookCommentReq) (res *v1.ListBookCommentRes, err error) {
	comments, count, err := service.Book().ListNewestComments(ctx, req.BookId)
	if err != nil {
		return nil, err
	}
	var listComments []v1.Comments
	for _, v := range comments {
		listComments = append(listComments, v1.Comments{
			Id:               v.Id,
			CommentContent:   v.CommentContent,
			CommentUser:      v.CommentUser,
			CommentUserId:    v.CommentUserId,
			CommentUserPhoto: v.CommentUserPhoto,
			CommentTime:      v.CommentTime,
		})
	}

	return &v1.ListBookCommentRes{
		CommentTotal: count,
		CommentArray: listComments,
	}, err
}

func (c *BookControllerV1) ListBookCategory(ctx context.Context, req *v1.ListBookCategoryReq) (res []*v1.ListBookCategoryRes, err error) {

	categoryArray, err := service.Book().FindWorkDirectionCategory(ctx, req.WorkDirection)
	if err != nil {
		return nil, err
	}
	var categoryInfo []*v1.ListBookCategoryRes
	for _, category := range categoryArray {
		categoryInfo = append(categoryInfo, &v1.ListBookCategoryRes{
			Id:   int(category.Id),
			Name: category.Name,
		})
	}
	return categoryInfo, nil
}

func (c *BookControllerV1) GetBookNewestRank(ctx context.Context, req *v1.GetBookNewestRankReq) (res []*v1.GetBookNewestRankRes, err error) {
	rank, err := service.Book().ListNewestRank(ctx)
	if err != nil {
		return nil, err
	}

	var bookNewestRank []*v1.GetBookNewestRankRes
	for _, v := range rank {
		bookNewestRankItem := &v1.GetBookNewestRankRes{
			BookRankInfo: v1.BookRankInfo{
				Id:                    strconv.FormatUint(v.Id, 10),
				CategoryId:            v.CategoryId,
				CategoryName:          v.CategoryName,
				PicUrl:                v.PicUrl,
				BookName:              v.BookName,
				AuthorName:            v.AuthorName,
				BookDesc:              v.BookDesc,
				WordCount:             v.WordCount,
				LastChapterName:       v.LastChapterName,
				LastChapterUpdateTime: v.LastChapterUpdateTime,
			},
		}
		bookNewestRank = append(bookNewestRank, bookNewestRankItem)
	}

	return bookNewestRank, nil
}

func (c *BookControllerV1) GetNextChapterId(ctx context.Context, req *v1.GetNextChapterIdReq) (res *v1.GetNextChapterIdRes, err error) {
	nextId, err := service.Book().FindNextChapterId(ctx, req.ChapterId)
	if err != nil {
		return nil, err
	}
	return &v1.GetNextChapterIdRes{Id: strconv.Itoa(nextId)}, nil
}

func (c *BookControllerV1) GetPreChapterId(ctx context.Context, req *v1.GetPreChapterIdReq) (res *v1.GetPreChapterIdRes, err error) {
	id, err := service.Book().FindPreChapterId(ctx, req.ChapterId)
	if err != nil {
		return nil, err
	}
	return &v1.GetPreChapterIdRes{Id: strconv.Itoa(id)}, nil
}

func (c *BookControllerV1) ListRecommendedBooks(ctx context.Context, req *v1.ListRecommendedBooksReq) (res []*v1.ListRecommendedBooksRes, err error) {
	reList, err := service.Book().ListRecBooks(ctx, req.BookId)
	if err != nil {
		return nil, err
	}
	var recListRes []*v1.ListRecommendedBooksRes
	for _, v := range reList {
		bookInfoRes := &v1.ListRecommendedBooksRes{
			BookInfo: v1.BookInfo{
				Id:              strconv.FormatUint(v.Id, 10),
				CategoryId:      v.CategoryId,
				CategoryName:    v.CategoryName,
				PicUrl:          v.PicUrl,
				BookName:        v.BookName,
				AuthorId:        v.AuthorId,
				AuthorName:      v.AuthorName,
				BookDesc:        v.BookDesc,
				BookStatus:      v.BookStatus,
				VisitCount:      v.VisitCount,
				WordCount:       v.WordCount,
				CommentCount:    v.CommentCount,
				FirstChapterId:  "1",
				LastChapterId:   strconv.FormatUint(v.LastChapterId, 10),
				LastChapterName: v.LastChapterName,
				UpdateTime:      v.UpdateTime,
			},
		}
		recListRes = append(recListRes, bookInfoRes)
	}
	return recListRes, nil
}

func (c *BookControllerV1) GetBookUpdateRank(ctx context.Context, req *v1.GetBookUpdateRankReq) (res []*v1.GetBookUpdateRankRes, err error) {
	rank, err := service.Book().ListUpdateRank(ctx)
	if err != nil {
		return nil, err
	}

	var bookUpdateRank []*v1.GetBookUpdateRankRes
	for _, item := range rank {
		resItem := &v1.GetBookUpdateRankRes{
			BookRankInfo: v1.BookRankInfo{
				Id:                    strconv.FormatUint(item.Id, 10),
				CategoryId:            item.CategoryId,
				CategoryName:          item.CategoryName,
				PicUrl:                item.PicUrl,
				BookName:              item.BookName,
				AuthorName:            item.AuthorName,
				BookDesc:              item.BookDesc,
				WordCount:             item.WordCount,
				LastChapterName:       item.LastChapterName,
				LastChapterUpdateTime: item.LastChapterUpdateTime,
			},
		}
		bookUpdateRank = append(bookUpdateRank, resItem)
	}

	return bookUpdateRank, nil
}

func (c *BookControllerV1) AddBookVisit(ctx context.Context, req *v1.AddBookVisitReq) (res *v1.AddBookVisitRes, err error) {
	err = service.Book().AddBookVisit(ctx, req.BookId)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *BookControllerV1) GetBookVisitRank(ctx context.Context, req *v1.GetBookVisitRankReq) (res []*v1.GetBookVisitRankRes, err error) {
	rank, err := service.Book().ListVisitRank(ctx)
	if err != nil {
		return nil, err
	}

	res = make([]*v1.GetBookVisitRankRes, 0, len(rank))
	for _, item := range rank {
		bookVisitRankRes := &v1.GetBookVisitRankRes{
			BookRankInfo: v1.BookRankInfo{
				Id:                    strconv.FormatUint(item.Id, 10),
				CategoryId:            item.CategoryId,
				CategoryName:          item.CategoryName,
				PicUrl:                item.PicUrl,
				BookName:              item.BookName,
				AuthorName:            item.AuthorName,
				BookDesc:              item.BookDesc,
				WordCount:             item.WordCount,
				LastChapterName:       item.LastChapterName,
				LastChapterUpdateTime: item.LastChapterUpdateTime,
			},
		}
		res = append(res, bookVisitRankRes)
	}
	return res, nil
}
