package book

import (
	"Literary_Snowlands/internal/consts"
	"Literary_Snowlands/internal/dao"
	"Literary_Snowlands/internal/model"
	"Literary_Snowlands/internal/model/entity"
	"Literary_Snowlands/internal/service"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
)

type sBook struct {
}

func init() {
	service.RegisterBook(New())
}

func New() service.IBook {
	return &sBook{}
}

func (s *sBook) AddBookVisit(ctx context.Context, bookId int) (err error) {
	_, err = dao.BookInfo.Ctx(ctx).Where("id = ?", bookId).Increment("visit_count", 1)
	if err != nil {
		return fmt.Errorf("addBookVisit wrong %s", err)
	}
	return nil
}

func (s *sBook) FindWorkDirectionCategory(ctx context.Context, workDirection int) (category []entity.BookCategory, err error) {
	var bookCategory []entity.BookCategory
	err = dao.BookCategory.Ctx(ctx).Where("work_direction = ?", workDirection).Scan(&bookCategory)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		} else {
			return nil, fmt.Errorf("FindWorkDirectionCategory %s", err)
		}
	}

	return bookCategory, err
}

func (s *sBook) FindBookInfoAndChapterNumById(ctx context.Context, id int) (info *entity.BookInfo, firstChapterID int, err error) {
	var bookInfo entity.BookInfo
	err = dao.BookInfo.Ctx(ctx).Where("id = ?", id).Scan(&bookInfo)
	if err != nil {
		return nil, 0, fmt.Errorf("FindBookInfoAndChapterNumById %s", err)
	}

	firstChapterInfo, err := dao.BookChapter.Ctx(ctx).Where("book_id = ?", id).OrderAsc("chapter_num").One()
	if err != nil {
		return nil, 0, fmt.Errorf("FindBookInfoAndChapterNumById %s", err)
	}
	bookChapterNumMap := firstChapterInfo.Map()
	firstChapterID = int(bookChapterNumMap["id"].(uint64))
	return &bookInfo, firstChapterID, nil
}

func (s *sBook) GetLastChapterAbout(ctx context.Context, id int) (chapterInfo *entity.BookChapter, total int, content string, err error) {

	lastChapterId, err := dao.BookInfo.Ctx(ctx).Where("id = ?", id).Value("last_chapter_id")
	if err != nil {
		return nil, 0, "", fmt.Errorf("GetLastChapterAbout %s", err)
	}

	err = dao.BookChapter.Ctx(ctx).Where("id = ?", lastChapterId).Scan(&chapterInfo)
	if chapterInfo == nil {
		return nil, 0, "", nil
	}

	total, err = dao.BookChapter.Ctx(ctx).Where("book_id = ?", id).Count()
	if err != nil {
		return nil, 0, "", fmt.Errorf("GetLastChapterAbout %s", err)
	}

	contentValue, err := dao.BookContent.Ctx(ctx).Where("chapter_id = ?", chapterInfo.Id).Value("content")
	if err != nil {
		return nil, 0, "", fmt.Errorf("GetLastChapterAbout %s", err)
	}

	content = contentValue.String()
	contentRune := []rune(content)
	if len(contentRune) > 30 {
		content = string(contentRune[0:100])
	}
	return chapterInfo, total, content, nil
}

func (s *sBook) ListRecBooks(ctx context.Context, id int) (re []*entity.BookInfo, err error) {
	var bookInfo entity.BookInfo
	err = dao.BookInfo.Ctx(ctx).Where("id = ?", id).Scan(&bookInfo)
	if err != nil {
		return nil, fmt.Errorf("ListRecBooks %s", err)
	}

	var recBookInfo []*entity.BookInfo
	err = dao.BookInfo.Ctx(ctx).Where("category_id = ?", bookInfo.CategoryId).
		Where("word_count > 0").
		OrderDesc("update_time").
		Limit(500).Scan(&recBookInfo)
	if err != nil {
		return nil, fmt.Errorf("ListRecBooks %s", err)
	}

	selected := make(map[int]struct{})
	var selectedBooks []*entity.BookInfo

	for i := 0; i < consts.REC_BOOK_COUNT && i < len(recBookInfo); i++ {
		for {
			randomIndex := rand.Intn(len(recBookInfo))
			if _, exists := selected[randomIndex]; !exists {
				selected[randomIndex] = struct{}{}
				selectedBooks = append(selectedBooks, recBookInfo[randomIndex])
				break
			}
			if len(selected) == len(recBookInfo) {
				break
			}
		}
	}

	return selectedBooks, nil
}

func (s *sBook) FindChapterList(ctx context.Context, id int) (chapters []*entity.BookChapter, err error) {
	err = dao.BookChapter.Ctx(ctx).Where("book_id = ?", id).OrderAsc("chapter_num").Scan(&chapters)
	if err != nil {
		return nil, fmt.Errorf("FindChapterList %s", err)
	}
	return chapters, nil
}

func (s *sBook) FindBookContentByChapterId(ctx context.Context, id int) (chapterInfo *entity.BookChapter, content string, bookInfo *entity.BookInfo, firstChapterId int, err error) {
	err = dao.BookChapter.Ctx(ctx).Where("id= ?", id).Scan(&chapterInfo)
	if err != nil {
		return nil, "", nil, 0, fmt.Errorf("FindBookContentByChapterId %s", err)
	}

	firstValue, err := dao.BookChapter.Ctx(ctx).Where("book_id = ?", chapterInfo.BookId).OrderAsc("chapter_num").Value("id")
	if err != nil {
		return nil, "", nil, 1, fmt.Errorf("FindBookContentByChapterId %s", err)
	}

	if firstValue != nil { // 检查 firstValue 是否为 nil
		firstChapterId = firstValue.Int()
	} else {
		return nil, "", nil, 0, nil
	}
	contentValue, err := dao.BookContent.Ctx(ctx).Where("chapter_id = ?", id).Value("content")
	content = contentValue.String()
	if err != nil {
		return nil, "", nil, 0, fmt.Errorf("FindBookContentByChapterId %s", err)
	}
	err = dao.BookInfo.Ctx(ctx).Where("id = ?", chapterInfo.BookId).Scan(&bookInfo)
	if err != nil {
		return nil, "", nil, 0, fmt.Errorf("FindBookContentByChapterId %s", err)
	}
	return chapterInfo, content, bookInfo, firstChapterId, err
}

func (s *sBook) FindPreChapterId(ctx context.Context, id int) (preId int, err error) {
	var chapterInfo entity.BookChapter
	err = dao.BookChapter.Ctx(ctx).Where("id = ?", id).Scan(&chapterInfo)
	if err != nil {
		return 0, fmt.Errorf("FindPreChapterId %s", err)
	}

	one, err := dao.BookChapter.Ctx(ctx).Where("book_id = ?", chapterInfo.BookId).
		Where("chapter_num", chapterInfo.ChapterNum-1).One()
	if err != nil {
		return 0, fmt.Errorf("FindPreChapterId %s", err)
	}
	preId = one["id"].Int()
	return preId, nil
}

func (s *sBook) FindNextChapterId(ctx context.Context, id int) (nextId int, err error) {
	var chapterInfo entity.BookChapter
	err = dao.BookChapter.Ctx(ctx).Where("id = ?", id).Scan(&chapterInfo)
	if err != nil {
		return 0, fmt.Errorf("FindNextChapterId %s", err)
	}
	one, err := dao.BookChapter.Ctx(ctx).Where("book_id = ?", chapterInfo.BookId).
		Where("chapter_num", chapterInfo.ChapterNum+1).One()
	if err != nil {
		return 0, fmt.Errorf("FindNextChapterId %s", err)
	}
	nextId = one["id"].Int()
	return nextId, err
}

func (s *sBook) ListVisitRank(ctx context.Context) (rank []*entity.BookInfo, err error) {
	var bookVisitRank []*entity.BookInfo
	err = dao.BookInfo.Ctx(ctx).Where("word_count > ?", 0).OrderDesc("visit_count").Limit(30).Scan(&bookVisitRank)
	if err != nil {
		return nil, fmt.Errorf("ListVisitRank %s", err)
	}
	return bookVisitRank, nil
}

func (s *sBook) ListNewestRank(ctx context.Context) (rank []*entity.BookInfo, err error) {
	var bookNewestRank []*entity.BookInfo
	err = dao.BookInfo.Ctx(ctx).Where("word_count > ?", 0).OrderDesc("create_time").Limit(30).Scan(&bookNewestRank)
	if err != nil {
		return nil, fmt.Errorf("ListNewestRank %s", err)
	}
	return bookNewestRank, nil
}

func (s *sBook) ListUpdateRank(ctx context.Context) (rank []*entity.BookInfo, err error) {
	var bookNewestRank []*entity.BookInfo
	err = dao.BookInfo.Ctx(ctx).Where("word_count > ?", 0).OrderDesc("update_time").Limit(30).Scan(&bookNewestRank)
	if err != nil {
		return nil, fmt.Errorf("ListUpdateRank %s", err)
	}
	return bookNewestRank, nil
}

func (s *sBook) ListNewestComments(ctx context.Context, bookId int) (comments []*model.Comment, commentCount int, err error) {

	count, err := dao.BookComment.Ctx(ctx).Where("book_id = ?", bookId).Count()
	if err != nil {
		return nil, 0, fmt.Errorf("ListNewestComments %s", err)
	}
	if count == 0 {
		return nil, 0, nil
	}

	// 查询最新的5条评论
	var doBookComments []entity.BookComment

	var commentRes []*model.Comment
	err = dao.BookComment.Ctx(ctx).Where("book_id = ?", bookId).OrderDesc("create_time").Limit(5).Scan(&doBookComments)
	if err != nil {
		return nil, 0, fmt.Errorf("ListNewestComments %s", err)
	}
	for _, v := range doBookComments {
		var userInfo entity.UserInfo
		dao.UserInfo.Ctx(ctx).Where("username = ?", v.UserId).Scan(&userInfo)
		commentRes = append(commentRes, &model.Comment{
			Id:               v.Id,
			CommentContent:   v.CommentContent,
			CommentUser:      userInfo.Username,
			CommentUserId:    userInfo.Id,
			CommentUserPhoto: userInfo.UserPhoto,
			CommentTime:      v.CreateTime,
		})
	}

	return commentRes, count, nil
}
