package news

import (
	"Literary_Snowlands/internal/model"
	"Literary_Snowlands/internal/service"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
)

type sNews struct {
}

func init() {
	service.RegisterNews(New())
}

func New() service.INews {
	return &sNews{}
}

func (s *sNews) FindLatestTwoNewsInfo() (newsInfo []model.NewsInfo, err error) {
	err = g.DB().Model("news_info ni").
		LeftJoin("news_content nc", "ni.id=nc.news_id").
		Fields("ni.id, ni.category_id, ni.category_name, ni.title, ni.source_name, ni.update_time, nc.content").
		Order("ni.create_time desc").
		Limit(2).
		Scan(&newsInfo)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		}
		return nil, err
	}
	return newsInfo, err
}

func (s *sNews) FindNewsInfoById(id int) (newsInfo *model.NewsInfo, err error) {
	err = g.DB().Model("news_info ni").
		LeftJoin("news_content nc", "ni.id = nc.news_id").
		Fields("ni.id,ni.category_id ,ni.category_name,ni.title,ni.source_name,ni.update_time,nc.content").
		Where("ni.id = ?", id).
		Scan(&newsInfo)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return newsInfo, err
}
