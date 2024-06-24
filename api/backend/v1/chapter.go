package v1

import "github.com/gogf/gf/v2/frame/g"

type GetChapterIdReq struct {
	g.Meta `path:"/api/author/book/chapter/{chapterId}" method:"get" summary:"添加文章"`
}

type GetChapterIdRes struct {
	ChapterName    string `json:"chapterName"`
	ChapterContent string `json:"chapterContent"`
	IsVip          string `json:"isVip"`
}
