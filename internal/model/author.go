package model

type AuthorRegister struct {
	UserId        int    `json:"userId"`
	PenName       string `json:"penName"`
	TelPhone      string `json:"telPhone"`
	ChatAccount   string `json:"chatAccount"`
	Email         string `json:"email"`
	WorkDirection int    `json:"workDirection"`
}

type AuthorBookInput struct {
	BookId         string `json:"bookId"`
	ChapterName    string `json:"chapterName"`
	ChapterContent string `json:"chapterContent"`
	IsVip          int    `json:"isVip"`
}

type BookChapter struct {
	ChapterName    string `json:"chapterName"`
	ChapterContent string `json:"chapterContent"`
	IsVip          int    `json:"isVip"`
}

type ListBook struct {
	BookId   string `json:"bookId"`
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
}
