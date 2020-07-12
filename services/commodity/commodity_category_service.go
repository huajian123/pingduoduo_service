package commodity_service

import (
	commodity_dao "pingduoduo_service/dao/commodity"
	"pingduoduo_service/datasource"
	commodity_model "pingduoduo_service/models/commodity"
)

type CommodityCategoryService interface {
	GetCommodityCategoryList(pageSize int, pageNum int) []commodity_model.CommodityCategory
	GetCommodityCategoryCount() int
	AddCommodityCategory(category commodity_model.CommodityCategory) error
}

type commodityCategoryService struct {
	dao *commodity_dao.CommodityCategoryDao
}

func (this *commodityCategoryService) GetCommodityCategoryList(pageSize int, pageNum int) []commodity_model.CommodityCategory {
	return this.dao.GetCommodityCategoryList(pageSize, pageNum)
}

func (this *commodityCategoryService) AddCommodityCategory(param commodity_model.CommodityCategory) error {
	return this.dao.AddCommodityCategoryList(param)
}

func (this *commodityCategoryService) GetCommodityCategoryCount() int {
	return this.dao.GetSearchListCount()
}

func NewCommodityCategoryService() CommodityCategoryService {
	return &commodityCategoryService{
		dao: commodity_dao.NewCommodityCategoryDao(datasource.GetDb()),
	}
}
