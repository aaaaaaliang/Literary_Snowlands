package search

import (
	"Literary_Snowlands/internal/dao"
	"Literary_Snowlands/internal/model"
	"Literary_Snowlands/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type sSearch struct{}

func init() {
	service.RegisterSearch(New())
}

func New() service.ISearch {
	return &sSearch{}
}

func (s *sSearch) SearchBooks(ctx context.Context, query model.SearchBooksQuery) (*model.SearchBooksInfo, error) {
	m := g.DB().Model("book_info").Ctx(ctx)
	var total int
	if query.Keyword != "" {
		var booksInfoList []model.ListInfo
		queryString := "%" + query.Keyword + "%"
		err := m.WhereLike("book_name", queryString).WhereOrLike("author_name", queryString).
			ScanAndCount(&booksInfoList, &total, true)
		if err != nil {
			return nil, err
		}
		for _, v := range booksInfoList {
			firstChapterId, err := dao.BookChapter.Ctx(ctx).Where("book_id", v.Id).
				OrderAsc("chapter_num").Value("id")
			if err != nil {
				return nil, err
			}
			v.FirstChapterId = firstChapterId.Int()
		}
		pages := total / query.PageSize

		searchBookInfo := &model.SearchBooksInfo{
			PageNum:  query.PageNum,
			PageSize: query.PageSize,
			Total:    total,
			List:     booksInfoList,
			Pages:    pages,
		}

		return searchBookInfo, nil

	}

	if query.WorkDirection != "" {
		m = m.Where("work_direction = ?", query.WorkDirection)
	}

	if query.BookStatus != "" {
		m = m.Where("book_status = ?", query.BookStatus)
	}

	if query.CategoryId != "" {
		m = m.Where("category_id = ?", query.CategoryId)
	}

	if query.WordCountMin != "" {
		m = m.Where("word_count > ?", query.WordCountMin)
	}

	if query.WordCountMax != "" {
		m = m.Where("word_count < ?", query.WordCountMax)
	}

	if query.UpdateTimeMin != "" {
		m = m.Where("update_time > ?", query.UpdateTimeMin)
	}

	if query.IsVip != "" {
		m = m.Where("update_time > ?", query.IsVip)
	}

	if query.Sort != "" {
		m = m.Order(query.Sort)
	}
	var booksInfoList []model.ListInfo
	err := m.ScanAndCount(&booksInfoList, &total, true)
	if err != nil {
		return nil, err
	}
	for _, v := range booksInfoList {
		firstChapterId, err := dao.BookChapter.Ctx(ctx).Where("book_id", v.Id).
			OrderAsc("chapter_num").Value("id")
		if err != nil {
			return nil, err
		}
		v.FirstChapterId = firstChapterId.Int()
	}
	pages := (total + query.PageSize - 1) / query.PageSize

	searchBookInfo := &model.SearchBooksInfo{
		PageNum:  query.PageNum,
		PageSize: query.PageSize,
		Total:    total,
		List:     booksInfoList,
		Pages:    pages,
	}

	return searchBookInfo, nil

}
