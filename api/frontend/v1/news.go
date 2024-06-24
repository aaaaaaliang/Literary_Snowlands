package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type QueryNewsInfoReq struct {
	g.Meta `path:"/api/front/news/{id}" method:"get" tags:"新闻模块" sm:"新闻列表查询"`
	Id     int `json:"id"`
}

type QueryNewsInfoRes struct {
	NewsInfo
}

type QueryNewsLatestInfoReq struct {
	g.Meta `path:"/api/front/news/latest_list" method:"get" tags:"新闻模块" sm:"最新新闻列表查询"`
}

type QueryNewsLatestInfoRes struct {
	NewsInfo
}

type NewsInfo struct {
	Id           int       `json:"id"`
	CategoryId   int       `json:"categoryId"`
	CategoryName string    `json:"categoryName"`
	SourceName   string    `json:"sourceName"`
	Title        string    `json:"title"`
	UpdateTime   time.Time `json:"updateTime"`
	Content      string    `json:"content"`
}
