package commodity_service

import (
	commodity_dao "pingduoduo_service/dao/commodity"
	"pingduoduo_service/datasource"
	commodity_model "pingduoduo_service/models/commodity"
)

type CommodityCategoryService interface {
	GetCommodityCategoryList(pageSize int, pageNum int) []commodity_model.CommodityCategory
	GetCommodityCategoryCount() int
}

type commodityCategoryService struct {
	dao *commodity_dao.CommodityCategoryDao
}

func (this *commodityCategoryService) GetCommodityCategoryList(pageSize int, pageNum int) []commodity_model.CommodityCategory {
	return this.dao.GetCommodityCategoryList(pageSize, pageNum)
}

func (this *commodityCategoryService) GetCommodityCategoryCount() int {
	return this.dao.GetSearchListCount()
}

func NewCommodityCategoryService() CommodityCategoryService {
	return &commodityCategoryService{
		dao: commodity_dao.NewCommodityCategoryDao(datasource.GetDb()),
	}
}
