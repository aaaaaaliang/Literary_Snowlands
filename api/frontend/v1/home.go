package v1

import "github.com/gogf/gf/v2/frame/g"

type QueryRecommendedBooksReq struct {
	g.Meta `path:"/api/front/home/books" method:"get" tags:"首页模块" sm:"首页小说推荐查询接口"`
}

type QueryRecommendedBooksRes struct {
	Type       int    `json:"type"`
	BookId     string `json:"bookId"`
	PicUrl     string `json:"picUrl"`
	BookName   string `json:"bookName"`
	AuthorName string `json:"authorName"`
	BookDesc   string `json:"bookDesc"`
}

type QueryFriendLinkBooksReq struct {
	g.Meta `path:"/api/front/home/friend_Link/list" method:"get" tags:"首页模块" sm:"友情链接列表查询"`
}

type QueryFriendLinkBooksRes struct {
	g.Meta   `path:"/api/front/home/friend_Link"`
	LinkName string `json:"linkName"`
	LinkUrl  string `json:"linkUrl"`
}
