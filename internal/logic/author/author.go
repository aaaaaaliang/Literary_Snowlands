package author

import (
	"Literary_Snowlands/internal/dao"
	"Literary_Snowlands/internal/model"
	"Literary_Snowlands/internal/model/do"
	"Literary_Snowlands/internal/model/entity"
	"Literary_Snowlands/internal/service"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"strconv"
)

type sAuthor struct {
}

func init() {
	service.RegisterAuthor(New())
}

func New() service.IAuthor {
	return &sAuthor{}
}

func (s *sAuthor) RegisterAuthor(ctx context.Context, registerInfo do.AuthorInfo) error {
	uid := ctx.Value("userId")
	registerInfo.UserId = uid
	one, err := dao.AuthorInfo.Ctx(ctx).Where("user_id = ?", uid).One()
	if err != nil {
		return err
	}
	if !one.IsEmpty() {
		return fmt.Errorf("该用户已注册为作家")
	}
	registerInfo.CreateTime = gtime.Now()
	registerInfo.UpdateTime = gtime.Now()
	if registerInfo.InviteCode == nil {
		registerInfo.InviteCode = 0
	}
	err = dao.AuthorInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.AuthorInfo.Ctx(ctx).Data(registerInfo).Insert()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *sAuthor) QueryAuthorStatus(ctx context.Context) (*int, error) {
	uid := ctx.Value("userId")
	var authorInfo entity.AuthorInfo
	err := dao.AuthorInfo.Ctx(ctx).Where("user_id = ?", uid).Scan(&authorInfo)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	status := int(authorInfo.Status)
	return &status, nil
}

func (s *sAuthor) CreateBook(ctx context.Context, bookInfo do.BookInfo) error {
	count, err := dao.BookInfo.Ctx(ctx).Where("book_name = ?", bookInfo.BookName).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("该书名已存在")
	}

	uid := ctx.Value("userId")
	var authorInfo do.AuthorInfo
	err = dao.AuthorInfo.Ctx(ctx).Where("user_id = ?", uid).Scan(&authorInfo)
	if err != nil {
		return err
	}
	bookInfo.Score = 0
	bookInfo.UpdateTime = gtime.Now()
	bookInfo.CreateTime = gtime.Now()
	bookInfo.AuthorId = authorInfo.Id
	bookInfo.AuthorName = authorInfo.PenName
	err = dao.BookInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.BookInfo.Ctx(ctx).Data(bookInfo).Insert()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *sAuthor) ListAuthorBookInfo(ctx context.Context, pageNum, pageSize int) (bookInfoList []entity.BookInfo, total int, err error) {
	uid := ctx.Value("userId")
	authorId, err := dao.AuthorInfo.Ctx(ctx).Where("user_id = ?", uid).Value("id")
	if err != nil {
		return nil, 0, err
	}
	if authorId.Int() == 0 {
		return nil, 0, nil
	}
	err = dao.BookInfo.Ctx(ctx).Page(pageNum, pageSize).Where("author_id = ?", authorId.Int()).ScanAndCount(&bookInfoList, &total, true)
	if err != nil {
		return nil, 0, err
	}

	return bookInfoList, total, nil
}

func (s *sAuthor) SaveBookChapter(ctx context.Context, input model.AuthorBookInput) error {
	err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var bookInfo entity.BookInfo
		err := dao.BookInfo.Ctx(ctx).Where("id = ?", input.BookId).Scan(&bookInfo)
		if err != nil {
			return err
		}
		if bookInfo.AuthorId != uint64(ctx.Value("userId").(int)) {
			return fmt.Errorf("当前小说与作者不匹配")
		}

		var bookChapterList []entity.BookChapter
		// 查询当前是否有章节
		err = dao.BookChapter.Ctx(ctx).
			Where("book_id = ?", input.BookId).
			OrderDesc("chapter_num").
			Scan(&bookChapterList)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				err = nil
			}
			return err
		}
		var latestBookChapterNum int
		if len(bookChapterList) == 0 {
			var bookChapter entity.BookChapter
			bookChapter.BookId, _ = strconv.ParseUint(input.BookId, 10, 64)
			bookChapter.ChapterNum = 1
			bookChapter.ChapterName = fmt.Sprintf("第%d章 %s", bookChapter.ChapterNum, input.ChapterName)
			bookChapter.WordCount = uint(len(input.ChapterContent))
			bookChapter.IsVip = uint(input.IsVip)
			bookChapter.CreateTime = gtime.Now()
			bookChapter.UpdateTime = gtime.Now()

			// 没问题  返回的的id是content的book_id
			insertId, err := dao.BookChapter.Ctx(ctx).Data(bookChapter).InsertAndGetId()
			if err != nil {
				return err
			}
			var bookChapterContent entity.BookContent
			bookChapterContent.ChapterId = uint64(insertId)
			bookChapterContent.Content = input.ChapterContent
			bookChapterContent.CreateTime = gtime.Now()
			bookChapterContent.UpdateTime = gtime.Now()
			_, err = dao.BookContent.Ctx(ctx).Data(bookChapterContent).Insert()
			if err != nil {
				return err
			}
			latestBookChapterNum = 1
		} else {
			latestBookChapterNum = int(bookChapterList[0].ChapterNum) + 1
			bookId, _ := strconv.ParseUint(input.BookId, 10, 64)
			bookChapterInfo := entity.BookChapter{
				BookId:      bookId,
				ChapterNum:  uint(latestBookChapterNum),
				ChapterName: input.ChapterName,
				WordCount:   uint(len(input.ChapterContent)),
				IsVip:       uint(input.IsVip),
				CreateTime:  gtime.Now(),
				UpdateTime:  gtime.Now(),
			}
			insertId, err := dao.BookChapter.Ctx(ctx).Data(bookChapterInfo).InsertAndGetId()
			if err != nil {
				return err
			}
			bookChapterContentInfo := entity.BookContent{
				ChapterId:  uint64(insertId),
				Content:    input.ChapterContent,
				CreateTime: gtime.Now(),
				UpdateTime: gtime.Now(),
			}
			_, err = dao.BookContent.Ctx(ctx).Data(bookChapterContentInfo).Insert()
			if err != nil {
				return err
			}
		}

		var updateBookInfo entity.BookInfo
		err = dao.BookInfo.Ctx(ctx).Where("id = ?", input.BookId).Scan(&updateBookInfo)
		if err != nil {
			return err
		}

		updateBookInfo.WordCount = updateBookInfo.WordCount + uint(len(input.ChapterContent))
		updateBookInfo.LastChapterId = uint64(latestBookChapterNum)
		updateBookInfo.LastChapterName = input.ChapterName
		updateBookInfo.UpdateTime = gtime.Now()
		updateBookInfo.LastChapterUpdateTime = gtime.Now()

		_, err = dao.BookInfo.Ctx(ctx).Data(g.Map{
			"word_count":               updateBookInfo.WordCount,
			"last_chapter_id":          updateBookInfo.LastChapterId,
			"last_chapter_name":        updateBookInfo.LastChapterName,
			"update_time":              updateBookInfo.UpdateTime,
			"last_chapter_update_time": updateBookInfo.LastChapterUpdateTime,
		}).Where("id = ?", updateBookInfo.Id).Update()

		if err != nil {
			return nil
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *sAuthor) DeleteChapterById(ctx context.Context, chapterId string) error {
	var bookChapterInfo entity.BookChapter
	err := dao.BookChapter.Ctx(ctx).Where("id = ?", chapterId).Scan(&bookChapterInfo)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		} else {
			return err
		}

	}
	if bookChapterInfo.BookId == 0 {
		bookChapterInfo.BookId = 1
	}
	var bookChapterInfoList []entity.BookChapter
	err = dao.BookChapter.Ctx(ctx).Where("book_id = ?", bookChapterInfo.BookId).OrderDesc("chapter_num").Scan(&bookChapterInfoList)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		} else {
			return err
		}
	}

	if bookChapterInfo.BookId == 0 {
		bookChapterInfo.BookId = 1
	}
	var bookInfo entity.BookInfo
	err = dao.BookInfo.Ctx(ctx).Where("id = ?", bookChapterInfo.BookId).Scan(&bookInfo)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		} else {
			return err
		}
	}
	if uint(bookInfo.LastChapterId) == bookChapterInfo.ChapterNum {
		_, err = dao.BookChapter.Ctx(ctx).Where("id = ?", chapterId).Delete()
		if err != nil {
			return err
		}
		_, err = dao.BookContent.Ctx(ctx).Where("chapter_id = ?", bookChapterInfo.Id).Delete()
		if err != nil {
			return err
		}

		bookInfo.WordCount = bookInfo.WordCount - bookChapterInfo.WordCount
		if len(bookChapterInfoList) == 1 {
			bookInfo.LastChapterId = 0
			bookInfo.LastChapterName = ""
			bookInfo.LastChapterUpdateTime = gtime.Now()
		} else {
			bookInfo.LastChapterId = uint64(bookChapterInfoList[1].ChapterNum)
			bookInfo.LastChapterName = bookChapterInfoList[1].ChapterName
			bookInfo.LastChapterUpdateTime = bookChapterInfoList[1].CreateTime
		}
		bookInfo.UpdateTime = gtime.Now()

		_, err = dao.BookInfo.Ctx(ctx).Data(g.Map{
			"word_count":               bookInfo.WordCount,
			"last_chapter_id":          bookInfo.LastChapterId,
			"last_chapter_name":        bookInfo.LastChapterName,
			"update_time":              gtime.Now(),
			"last_chapter_update_time": bookInfo.LastChapterUpdateTime,
		}).Where("id = ?", bookInfo.Id).Update()
		if err != nil {
			return err
		}
		return nil
	} else {
		var bookInfoDto entity.BookInfo
		bookInfoDto.WordCount = bookInfo.WordCount - bookChapterInfo.WordCount
		bookInfoDto.UpdateTime = gtime.Now()
		_, err = dao.BookChapter.Ctx(ctx).Where("id = ?", chapterId).Delete()
		if err != nil {
			return err
		}
		_, err = dao.BookContent.Ctx(ctx).Where("chapter_id = ?", bookChapterInfo.Id).Delete()
		if err != nil {
			return err
		}

		bookInfo.WordCount = bookInfo.WordCount - bookChapterInfo.WordCount

		_, err = dao.BookInfo.Ctx(ctx).Data(g.Map{
			"word_count":  bookInfo.WordCount,
			"update_time": gtime.Now(),
		}).Where("id = ?", bookInfo.Id).Update()

		return nil
	}
}

func (s *sAuthor) GetChapterById(ctx context.Context, chapterId string) (*model.BookChapter, error) {
	var chapterInfo entity.BookChapter
	err := dao.BookChapter.Ctx(ctx).Where("id = ?", chapterId).Scan(&chapterInfo)
	if err != nil {
		return nil, err
	}

	var bookContent entity.BookContent
	err = dao.BookContent.Ctx(ctx).Where("chapter_id = ?", chapterId).Scan(&bookContent)
	if err != nil {
		return nil, err
	}

	// 在这里初始化 chapter
	chapter := &model.BookChapter{
		ChapterName:    chapterInfo.ChapterName,
		IsVip:          int(chapterInfo.IsVip),
		ChapterContent: bookContent.Content,
	}

	return chapter, nil
}

func (s *sAuthor) UpdateChapterById(ctx context.Context, chapterId string, chapter model.BookChapter) error {
	var err error
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var bookChapterInfo entity.BookChapter
		err = dao.BookChapter.Ctx(ctx).Where("id = ?", chapterId).Scan(&bookChapterInfo)
		if err != nil {
			return err
		}
		var bookInfo entity.BookInfo
		err = dao.BookInfo.Ctx(ctx).Where("id = ?", bookChapterInfo.BookId).Scan(&bookInfo)
		if err != nil {
			return err
		}
		var beginWordCount = bookChapterInfo.WordCount

		bookChapterInfo.ChapterName = chapter.ChapterName
		bookChapterInfo.WordCount = uint(len(chapter.ChapterContent))
		bookChapterInfo.IsVip = uint(chapter.IsVip)
		bookChapterInfo.UpdateTime = gtime.Now()

		var bookContent entity.BookContent
		err = dao.BookContent.Ctx(ctx).Where("chapter_id = ?", chapterId).Scan(&bookContent)
		if err != nil {
			return err
		}
		bookContent.Content = chapter.ChapterContent
		bookContent.UpdateTime = gtime.Now()

		updateWordCount := uint(len(chapter.ChapterContent)) - beginWordCount
		bookInfo.WordCount = bookInfo.WordCount + updateWordCount

		if bookChapterInfo.ChapterNum == uint(bookInfo.LastChapterId) {
			bookInfo.LastChapterName = chapter.ChapterName
			bookInfo.LastChapterUpdateTime = gtime.Now()
		}
		bookInfo.UpdateTime = gtime.Now()

		// 执行更新操作
		_, err = dao.BookInfo.Ctx(ctx).Data(g.Map{
			"word_count":               bookInfo.WordCount,
			"last_chapter_name":        bookInfo.LastChapterName,
			"last_chapter_update_time": bookInfo.LastChapterUpdateTime,
			"update_time":              bookInfo.UpdateTime,
		}).Where("id = ?", bookInfo.Id).Update()
		if err != nil {
			return err
		}

		_, err = dao.BookChapter.Ctx(ctx).Data(g.Map{
			"chapter_name": bookChapterInfo.ChapterName,
			"word_count":   bookChapterInfo.WordCount,
			"is_vip":       bookChapterInfo.IsVip,
			"update_time":  bookChapterInfo.UpdateTime,
		}).Where("id = ?", chapterId).Update()
		if err != nil {
			return err
		}

		_, err = dao.BookContent.Ctx(ctx).Data(g.Map{
			"content":     bookContent.Content,
			"update_time": bookContent.UpdateTime,
		}).Where("chapter_id = ?", chapterId).Update()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *sAuthor) ListBookChapters(ctx context.Context, book model.ListBook) (listBookChapters []entity.BookChapter, total int, err error) {
	err = dao.BookChapter.Ctx(ctx).Page(book.PageNum, book.PageSize).
		Where("book_id = ?", book.BookId).
		OrderDesc("chapter_num").
		ScanAndCount(&listBookChapters, &total, true)
	if err != nil {
		return nil, 0, err
	}

	return listBookChapters, total, err
}
