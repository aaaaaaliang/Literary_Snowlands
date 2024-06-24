package frontend

import (
	"Literary_Snowlands/api/frontend"
	v1 "Literary_Snowlands/api/frontend/v1"
	"Literary_Snowlands/internal/model"
	"Literary_Snowlands/internal/service"
	"context"
	"strconv"
)

type SearchControllerV1 struct {
}

func NewSearchControllerV1() frontend.ISearchV1 {
	return &SearchControllerV1{}
}

func (c *SearchControllerV1) SearchBooks(ctx context.Context, req *v1.SearchBooksReq) (*v1.SearchBooksRes, error) {
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 构造搜索条件
	searchQuery := model.SearchBooksQuery{
		Keyword:       req.Keyword,
		WorkDirection: req.WorkDirection,
		CategoryId:    req.CategoryId,
		IsVip:         req.IsVip,
		BookStatus:    req.BookStatus,
		WordCountMin:  req.WordCountMin,
		WordCountMax:  req.WordCountMax,
		UpdateTimeMin: req.UpdateTimeMin,
		Sort:          req.Sort,
		PageNum:       req.PageNum,
		PageSize:      req.PageSize,
	}

	// 调用内部服务层的搜索方法
	searchResult, err := service.Search().SearchBooks(ctx, searchQuery)
	if err != nil {
		return nil, err
	}

	// 转换内部服务层返回的结果为API响应格式
	res := &v1.SearchBooksRes{
		PageNum:  searchResult.PageNum,
		PageSize: searchResult.PageSize,
		Total:    searchResult.Total,
		Pages:    searchResult.Pages,
		List:     make([]v1.ListInfo, 0, len(searchResult.List)),
	}

	for _, item := range searchResult.List {
		res.List = append(res.List, v1.ListInfo{
			Id:              strconv.Itoa(item.Id),
			CategoryId:      item.CategoryId,
			CategoryName:    item.CategoryName,
			PicUrl:          item.PicUrl,
			BookName:        item.BookName,
			AuthorId:        item.AuthorId,
			AuthorName:      item.AuthorName,
			BookDesc:        item.BookDesc,
			BookStatus:      item.BookStatus,
			VisitCount:      item.VisitCount,
			WordCount:       item.WordCount,
			CommentCount:    item.CommentCount,
			FirstChapterId:  item.FirstChapterId,
			LastChapterId:   item.LastChapterId,
			LastChapterName: item.LastChapterName,
			UpdateTime:      item.UpdateTime,
		})
	}

	return res, nil
}
