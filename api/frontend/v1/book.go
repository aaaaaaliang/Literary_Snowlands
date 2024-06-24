package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type AddBookVisitReq struct {
	g.Meta `path:"/api/front/book/visit" method:"post" tags:"小说模块" sm:"增加访问量"`
	BookId int `json:"bookId"`
}

type AddBookVisitRes struct {
}

type ListBookCategoryReq struct {
	g.Meta        `path:"/api/front/book/category/list" method:"get" tags:"小说模块" summary:"小说分类列表查询"`
	WorkDirection int `json:"workDirection"` // 0 男频  1 女频
}

type ListBookCategoryRes struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type GetBookInfoByIdReq struct {
	g.Meta `path:"/api/front/book/{id}" method:"get" tags:"小说模块" summary:"小说信息查询"`
	Id     int `json:"id"`
}

type GetBookInfoByIdRes struct {
	BookInfo
}

type BookInfo struct {
	Id              string      `json:"id"`
	CategoryId      uint64      `json:"categoryId"`
	CategoryName    string      `json:"categoryName"`
	PicUrl          string      `json:"picUrl"`
	BookName        string      `json:"bookName"`
	AuthorId        uint64      `json:"authorId"`
	AuthorName      string      `json:"authorName"`
	BookDesc        string      `json:"bookDesc"`
	BookStatus      uint        `json:"bookStatus"`
	VisitCount      uint64      `json:"visitCount"`
	WordCount       uint        `json:"wordCount"`
	CommentCount    uint        `json:"commentCount"`
	FirstChapterId  string      `json:"firstChapterId"`
	LastChapterId   string      `json:"lastChapterId"`
	LastChapterName string      `json:"lastChapterName"`
	UpdateTime      *gtime.Time `json:"updateTime"`
}
type GetLastChapterAboutReq struct {
	g.Meta `path:"/api/front/book/last_chapter/about" method:"get" tags:"小说模块" sm:"小说最新章节相关信息查询接口"`
	BookId int `json:"bookId"`
}

type GetLastChapterAboutRes struct {
	ChapterInfo    `json:"chapterInfo"`
	ChapterTotal   int    `json:"chapterTotal"`
	ContentSummary string `json:"contentSummary"`
}

type ChapterInfo struct {
	Id                string      `json:"id"`
	BookId            string      `json:"bookId"`
	ChapterNum        uint        `json:"chapterNum"`
	ChapterName       string      `json:"chapterName"`
	ChapterWordCount  uint        `json:"chapterWordCount"`
	ChapterUpdateTime *gtime.Time `json:"chapterUpdateTime"`
	IsVip             uint        `json:"isVip"`
}

type ListRecommendedBooksReq struct {
	g.Meta `path:"/api/front/book/rec_list" method:"get" tags:"小说模块" sm:"小说推荐列表查询接口"`
	BookId int `json:"bookId"`
}

type ListRecommendedBooksRes struct {
	BookInfo
}

type ListChaptersReq struct {
	g.Meta `path:"/api/front/book/chapter/list" method:"get" tags:"小说模块" sm:"小说章节列表查询接口"`
	BookId int `json:"bookId"`
}

type ListChaptersRes struct {
	Id                string      `json:"id"`
	BookId            string      `json:"bookId"`
	ChapterNum        uint        `json:"chapterNum"`
	ChapterName       string      `json:"chapterName"`
	ChapterWordCount  uint        `json:"chapterWordCount"`
	ChapterUpdateTime *gtime.Time `json:"chapterUpdateTime"`
	IsVip             uint        `json:"isVip"`
}

type GetBookContentAboutReq struct {
	g.Meta    `path:"/api/front/book/content/{chapterId}" method:"get" tags:"小说模块" sm:"小说内容相关信息查询"`
	ChapterId int `json:"chapterId"`
}

type GetBookContentAboutRes struct {
	BookInfo    `json:"bookInfo"`
	ChapterInfo `json:"chapterInfo"`
	BookContent string `json:"bookContent"`
}

type GetPreChapterIdReq struct {
	g.Meta    `path:"/api/front/book/pre_chapter_id/{chapterId}" tags:"小说模块" sm:"上一章"`
	ChapterId int `json:"chapterId"`
}

type GetPreChapterIdRes struct {
	Id string `json:"id"`
}

type GetNextChapterIdReq struct {
	g.Meta    `path:"/api/front/book/next_chapter_id/{chapterId}" tags:"小说模块" sm:"下一章"`
	ChapterId int `json:"chapterId"`
}

type GetNextChapterIdRes struct {
	Id string `json:"id"`
}

type BookRankInfo struct {
	Id                    string      `json:"id"`
	CategoryId            uint64      `json:"categoryId"`
	CategoryName          string      `json:"categoryName"`
	PicUrl                string      `json:"picUrl"`
	BookName              string      `json:"bookName"`
	AuthorName            string      `json:"authorName"`
	BookDesc              string      `json:"bookDesc"`
	WordCount             uint        `json:"wordCount"`
	LastChapterName       string      `json:"lastChapterName"`
	LastChapterUpdateTime *gtime.Time `json:"lastChapterUpdateTime"`
}

type GetBookVisitRankReq struct {
	g.Meta `path:"/api/front/book/visit_rank" method:"get" tags:"小说模块" sum:"小说点击榜查询接口"`
}
type GetBookVisitRankRes struct {
	BookRankInfo
}

type GetBookNewestRankReq struct {
	g.Meta `path:"/api/front/book/newest_rank" tags:"小说模块" sum:"小说创建榜查询接口"`
}
type GetBookNewestRankRes struct {
	BookRankInfo
}

type GetBookUpdateRankReq struct {
	g.Meta `path:"/api/front/book/update_rank" tags:"小说模块" sum:"小说更新榜查询接口"`
}

type GetBookUpdateRankRes struct {
	BookRankInfo
}

type ListBookCommentReq struct {
	g.Meta `path:"/api/front/book/comment/newest_list" tags:"小说模块" sum:"小说最新评论查询接口"`
	BookId int `json:"bookId"`
}

type ListBookCommentRes struct {
	CommentTotal int `json:"commentTotal"`
	CommentArray `json:"comments"`
}
type CommentArray []Comments

type Comments struct {
	Id               uint64      `json:"id"`
	CommentContent   string      `json:"commentContent"`
	CommentUser      string      `json:"commentUser"`
	CommentUserId    uint64      `json:"commentUserId"`
	CommentUserPhoto string      `json:"commentUserPhoto"`
	CommentTime      *gtime.Time `json:"commentTime"`
}
