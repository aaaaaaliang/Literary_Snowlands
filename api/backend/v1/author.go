package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type AuthorRegisterReq struct {
	g.Meta        `path:"/api/author/register" method:"post" tags:"作家模块" sm:"作家信息注册"`
	PenName       string `json:"penName"`
	TelPhone      string `json:"telPhone"`
	ChatAccount   string `json:"chatAccount"`
	Email         string `json:"email"`
	WorkDirection int    `json:"workDirection"`
}

type AuthorRegisterRes struct {
}

type QueryAuthorStatusReq struct {
	g.Meta `path:"/api/author/status"    tags:"作家模块" sm:"作家状态查询"`
}

type QueryAuthorStatusRes struct {
	Status *int `json:"status"`
}

type PublishBookReq struct {
	g.Meta `path:"/api/author/book" method:"post" tags:"作家模块" sm:"小说发布接口"`
	BookInfo
}

type PublishBookRes struct {
}

type BookInfo struct {
	WorkDirection int    `json:"workDirection"`
	CategoryId    int    `json:"categoryId"`
	CategoryName  string `json:"categoryName"`
	PicUrl        string `json:"picUrl"`
	BookName      string `json:"bookName"`
	BookDesc      string `json:"bookDesc"`
	IsVip         int    `json:"isVip"`
}

type PublishBookListReq struct {
	g.Meta   `path:"/api/author/books" tags:"作家模块" sm:"小说发布列表查询"`
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

type PublishBookListRes struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
	List     `json:"list"`
	Pages    int `json:"pages"`
}
type List []BookList

type BookList struct {
	Id              string     `json:"id"`
	CategoryId      string     `json:"categoryId"`
	CategoryName    string     `json:"categoryName"`
	PicUrl          string     `json:"picUrl"`
	BookName        string     `json:"bookName"`
	AuthorId        int        `json:"authorId"`
	AuthorName      string     `json:"authorName"`
	BookDesc        string     `json:"bookDesc"`
	BookStatus      int        `json:"bookStatus"`
	VisitCount      int        `json:"visitCount"`
	WordCount       int        `json:"wordCount"`
	CommentCount    int        `json:"commentCount"`
	FirstChapterId  string     `json:"firstChapterId"`
	LastChapterId   string     `json:"lastChapterId"`
	LastChapterName string     `json:"lastChapterName"`
	UpdateTime      gtime.Time `json:"updateTime"`
}

type PublishBookChapterReq struct {
	g.Meta `path:"/api/author/book/chapter/{bookId}" method:"post" tags:"作家模块" sm:"小说章节发布"`

	BookId         string `json:"bookId"`
	ChapterName    string `json:"chapterName"`
	ChapterContent string `json:"chapterContent"`
	IsVip          int    `json:"isVip"`
}

type PublishBookChapterRes struct {
}

type DeleteBookChapterReq struct {
	g.Meta    `path:"/api/author/book/chapter/{chapterId}" method:"delete" tags:"作家模块" sm:"小说章节删除"`
	ChapterId string `json:"chapterId"`
}
type DeleteBookChapterRes struct {
}

type GetBookChapterReq struct {
	g.Meta    `path:"/api/author/book/chapter/{chapterId}" method:"get" tags:"小说章节查询"`
	ChapterId string `json:"chapterId"`
}
type GetBookChapterRes struct {
	ChapterName    string `json:"chapterName"`
	ChapterContent string `json:"chapterContent"`
	IsVip          int    `json:"isVip"`
}

type UpdateBookChapterReq struct {
	g.Meta         `path:"/api/author/book/chapter/{chapterId}" method:"put" tags:"作家模块" sm:"小说章节修改"`
	ChapterId      string `json:"chapterId"`
	ChapterName    string `json:"chapterName"`
	ChapterContent string `json:"chapterContent"`
	IsVip          int    `json:"isVip"`
}

type UpdateBookChapterRes struct {
}

type ListBookChaptersReq struct {
	g.Meta   `path:"/api/author/book/chapters/{bookId}" tags:"作家模块z" sm:"小说章节发布接口"`
	BookId   string `json:"bookId"`
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
}

type ListBookChaptersRes struct {
	PageNum     int `json:"pageNum"`
	PageSize    int `json:"pageSize"`
	Total       int `json:"total"`
	PublishList `json:"list"`
	Pages       int `json:"pages"`
}

type PublishList []PublishBook

type PublishBook struct {
	Id                string     `json:"id"`
	BookId            string     `json:"bookId"`
	ChapterNum        string     `json:"chapterNum"`
	ChapterName       string     `json:"chapterName"`
	ChapterWordCount  int        `json:"chapterWordCount"`
	ChapterUpdateTime gtime.Time `json:"chapterUpdateTime"`
	IsVip             int        `json:"isVip"`
}
