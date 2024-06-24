package model

import (
	"time"
)

type SearchBooksQuery struct {
	Keyword       string `json:"keyword"`
	WorkDirection string `json:"workDirection"`
	CategoryId    string `json:"categoryId"`
	IsVip         string `json:"isVip"`
	BookStatus    string `json:"bookStatus"`
	WordCountMin  string `json:"wordCountMin"`
	WordCountMax  string `json:"wordCountMax"`
	UpdateTimeMin string `json:"updateTimeMin"`
	Sort          string `json:"sort"`
	PageNum       int    `json:"pageNum"`
	PageSize      int    `json:"pageSize"`
}

type List []ListInfo
type SearchBooksInfo struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
	List     `json:"list"`
	Pages    int `json:"pages"`
}

type ListInfo struct {
	Id              int       `json:"id"`
	CategoryId      int       `json:"categoryId"`
	CategoryName    string    `json:"categoryName"`
	PicUrl          string    `json:"picUrl"`
	BookName        string    `json:"bookName"`
	AuthorId        int       `json:"authorId"`
	AuthorName      string    `json:"authorName"`
	BookDesc        string    `json:"bookDesc"`
	BookStatus      int       `json:"bookStatus"`
	VisitCount      int       `json:"visitCount"`
	WordCount       int       `json:"wordCount"`
	CommentCount    int       `json:"commentCount"`
	FirstChapterId  int       `json:"firstChapterId"`
	LastChapterId   int       `json:"lastChapterId"`
	LastChapterName string    `json:"lastChapterName"`
	UpdateTime      time.Time `json:"updateTime"`
}
