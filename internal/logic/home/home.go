package home

import (
	"Literary_Snowlands/internal/dao"
	"Literary_Snowlands/internal/model"
	"Literary_Snowlands/internal/model/entity"
	"Literary_Snowlands/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"log"
)

type sHome struct {
}

func init() {
	service.RegisterHome(New())
}

func New() service.IHome {
	return &sHome{}
}

func (s *sHome) FindRecommendBook(ctx context.Context) (recommendBooksInfo []model.RecommendedBooks, err error) {
	// 执行关联查询
	err = g.Model("home_book hb").Ctx(ctx).
		LeftJoin("book_info bi", "hb.book_id=bi.id").
		Fields("bi.author_name, bi.book_desc, bi.id as bookId, bi.book_name, bi.pic_url, hb.type").
		Scan(&recommendBooksInfo)
	if err != nil {
		return nil, err
	}
	return recommendBooksInfo, nil
}

func (s *sHome) QueryFriendLinkList(ctx context.Context) (homeLinkList []entity.HomeFriendLink, err error) {
	log.Println("执行了QueryFriendLinkList")
	err = dao.HomeFriendLink.Ctx(ctx).OrderAsc("sort").Scan(&homeLinkList)
	if err != nil {
		return nil, err
	}
	return homeLinkList, nil
}
