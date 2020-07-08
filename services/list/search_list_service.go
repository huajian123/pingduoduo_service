package list_service

import (
	list_dao "pingduoduo_service/dao/list"
	"pingduoduo_service/datasource"
	list_models "pingduoduo_service/models/list"
)

type SearchListService interface {
	GetSearchList() []list_models.SearchList
}

// 类
type searchListService struct {
	dao *list_dao.SearchListDao
}

func (s *searchListService)GetSearchList() []list_models.SearchList{
	return s.dao.GetSearchList()
}

func NewSearchListService()  SearchListService{
	return &searchListService{
		dao: list_dao.NewSearchListDao(datasource.GetDb()),
	}
}
