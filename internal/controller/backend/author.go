package backend

import (
	"Literary_Snowlands/api/backend"
	v1 "Literary_Snowlands/api/backend/v1"
	"Literary_Snowlands/internal/model"
	"Literary_Snowlands/internal/model/do"
	"Literary_Snowlands/internal/service"
	"context"
	"strconv"
)

type AuthorControllerV1 struct {
}

func NewAuthorControllerV1() backend.IAuthorV1 {
	return &AuthorControllerV1{}
}

func (c *AuthorControllerV1) DeleteChapter(ctx context.Context, req *v1.DeleteBookChapterReq) (res *v1.DeleteBookChapterRes, err error) {
	err = service.Author().DeleteChapterById(ctx, req.ChapterId)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *AuthorControllerV1) GetBookChapter(ctx context.Context, req *v1.GetBookChapterReq) (res *v1.GetBookChapterRes, err error) {
	chapterInfo, err := service.Author().GetChapterById(ctx, req.ChapterId)
	if err != nil {
		return nil, err
	}
	chapterRes := &v1.GetBookChapterRes{
		ChapterName:    chapterInfo.ChapterName,
		ChapterContent: chapterInfo.ChapterContent,
		IsVip:          chapterInfo.IsVip,
	}

	return chapterRes, nil
}

func (c *AuthorControllerV1) ListBookChapters(ctx context.Context, req *v1.ListBookChaptersReq) (res *v1.ListBookChaptersRes, err error) {
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	bookReq := model.ListBook{
		BookId:   req.BookId,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	}
	chapters, total, err := service.Author().ListBookChapters(ctx, bookReq)
	if err != nil {
		return nil, err
	}
	totalPages := total / req.PageSize
	if total%req.PageSize != 0 {
		totalPages += 1
	}

	list := make([]v1.PublishBook, len(chapters))
	for i, v := range chapters {
		list[i] = v1.PublishBook{
			Id:                strconv.FormatUint(v.Id, 10),
			BookId:            strconv.FormatUint(v.BookId, 10),
			ChapterNum:        strconv.Itoa(int(v.ChapterNum)),
			ChapterName:       v.ChapterName,
			ChapterWordCount:  int(v.WordCount),
			ChapterUpdateTime: *v.UpdateTime,
			IsVip:             int(v.IsVip),
		}
	}

	listBookRes := &v1.ListBookChaptersRes{
		PageNum:     req.PageNum,
		PageSize:    req.PageSize,
		Total:       total,
		PublishList: list,
		Pages:       totalPages,
	}

	return listBookRes, nil
}

func (c *AuthorControllerV1) PublishBook(ctx context.Context, req *v1.PublishBookReq) (res *v1.PublishBookRes, err error) {
	bookInfo := do.BookInfo{
		WorkDirection: req.WorkDirection,
		CategoryId:    req.CategoryId,
		CategoryName:  req.CategoryName,
		PicUrl:        req.PicUrl,
		BookName:      req.BookName,
		BookDesc:      req.BookDesc,
		IsVip:         req.IsVip,
	}

	err = service.Author().CreateBook(ctx, bookInfo)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *AuthorControllerV1) PublishBookChapter(ctx context.Context, req *v1.PublishBookChapterReq) (res *v1.PublishBookChapterRes, err error) {
	reqInfo := model.AuthorBookInput{
		BookId:         req.BookId,
		ChapterName:    req.ChapterName,
		ChapterContent: req.ChapterContent,
		IsVip:          req.IsVip,
	}
	err = service.Author().SaveBookChapter(ctx, reqInfo)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *AuthorControllerV1) PublishBookList(ctx context.Context, req *v1.PublishBookListReq) (*v1.PublishBookListRes, error) {

	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	booksInfoList, total, err := service.Author().ListAuthorBookInfo(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}

	totalPages := total / req.PageSize
	if total%req.PageSize != 0 {
		totalPages += 1
	}

	bookList := make([]v1.BookList, len(booksInfoList))
	for i, book := range booksInfoList {
		bookList[i] = v1.BookList{
			Id:              strconv.FormatUint(book.Id, 10),
			CategoryId:      strconv.FormatUint(book.CategoryId, 10),
			CategoryName:    book.CategoryName,
			PicUrl:          book.PicUrl,
			BookName:        book.BookName,
			AuthorId:        int(book.AuthorId),
			AuthorName:      book.AuthorName,
			BookDesc:        book.BookDesc,
			BookStatus:      int(book.BookStatus),
			VisitCount:      int(book.VisitCount),
			WordCount:       int(book.WordCount),
			CommentCount:    int(book.CommentCount),
			FirstChapterId:  strconv.Itoa(1),
			LastChapterId:   strconv.FormatUint(book.LastChapterId, 10),
			LastChapterName: book.LastChapterName,
			UpdateTime:      *book.UpdateTime,
		}
	}

	res := &v1.PublishBookListRes{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Total:    total,
		List:     bookList,
		Pages:    totalPages,
	}
	return res, nil
}

func (c *AuthorControllerV1) QueryAuthorStatus(ctx context.Context, req *v1.QueryAuthorStatusReq) (res *v1.QueryAuthorStatusRes, err error) {
	status, err := service.Author().QueryAuthorStatus(ctx)
	if err != nil {
		return nil, err
	}

	if status == nil {
		return &v1.QueryAuthorStatusRes{Status: nil}, nil
	} else {
		return &v1.QueryAuthorStatusRes{
			Status: status,
		}, nil
	}

}

func (c *AuthorControllerV1) AuthorRegister(ctx context.Context, req *v1.AuthorRegisterReq) (res *v1.AuthorRegisterRes, err error) {
	authorModelInfo := do.AuthorInfo{
		PenName:       req.PenName,
		TelPhone:      req.TelPhone,
		ChatAccount:   req.ChatAccount,
		Email:         req.Email,
		WorkDirection: uint(req.WorkDirection),
	}
	err = service.Author().RegisterAuthor(ctx, authorModelInfo)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *AuthorControllerV1) UpdateBookChapter(ctx context.Context, req *v1.UpdateBookChapterReq) (res *v1.UpdateBookChapterRes, err error) {
	var bookChapter model.BookChapter
	bookChapter.ChapterName = req.ChapterName
	bookChapter.ChapterContent = req.ChapterContent
	bookChapter.IsVip = req.IsVip
	err = service.Author().UpdateChapterById(ctx, req.ChapterId, bookChapter)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
