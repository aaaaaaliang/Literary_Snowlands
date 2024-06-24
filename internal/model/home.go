package model

type RecommendedBooks struct {
	Type       int    `json:"type"`
	BookId     string `json:"bookId"`
	PicUrl     string `json:"picUrl"`
	BookName   string `json:"bookName"`
	AuthorName string `json:"authorName"`
	BookDesc   string `json:"bookDesc"`
}
