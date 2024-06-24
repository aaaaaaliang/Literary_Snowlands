package model

import "time"

type NewsInfo struct {
	Id           int       `json:"id"`
	CategoryId   int       `json:"categoryId"`
	CategoryName string    `json:"categoryName"`
	SourceName   string    `json:"sourceName"`
	Title        string    `json:"title"`
	UpdateTime   time.Time `json:"updateTime"`
	Content      string    `json:"content"`
}
