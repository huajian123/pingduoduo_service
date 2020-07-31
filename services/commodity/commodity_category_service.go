package commodity_service

import (
	commodity_dao "pingduoduo_service/dao/commodity"
	"pingduoduo_service/datasource"
	commodity_model "pingduoduo_service/models/commodity"
)

type CommodityCategoryService interface {
	GetCommodityCategoryList(pageSize int, pageNum int, searchParam map[string]interface{}) []commodity_model.CommodityCategory
	GetCommodityCategoryCount(searchParam map[string]interface{}) int
	AddCommodityCategory(category commodity_model.CommodityCategory) error
	DelCommodityCategory(id int) error
	UpdateCommodityCategory(commodity_model.CommodityCategory) commodity_model.CommodityCategory
	GetCommodityCategoryDetail(id int) commodity_model.CommodityCategory
}

type commodityCategoryService struct {
	dao *commodity_dao.CommodityCategoryDao
}

func (this *commodityCategoryService) GetCommodityCategoryList(pageSize int, pageNum int, searchParam map[string]interface{}) []commodity_model.CommodityCategory {

	return this.dao.GetCommodityCategoryList(pageSize, pageNum, searchParam)
}

func (this *commodityCategoryService) AddCommodityCategory(param commodity_model.CommodityCategory) error {
	return this.dao.AddCommodityCategoryList(param)
}

func (this *commodityCategoryService) GetCommodityCategoryCount(searchParam map[string]interface{}) int {
	return this.dao.GetSearchListCount(searchParam)
}

func (this *commodityCategoryService) DelCommodityCategory(id int) error {
	return this.dao.DelCommodityCategory(id)
}

func (this *commodityCategoryService) GetCommodityCategoryDetail(id int) commodity_model.CommodityCategory {
	return this.dao.GetCommodityCategoryDetail(id)
}

func (this *commodityCategoryService) UpdateCommodityCategory(param commodity_model.CommodityCategory) commodity_model.CommodityCategory {
	return this.dao.UpdateCommodityCategory(param)
}

func NewCommodityCategoryService() CommodityCategoryService {
	return &commodityCategoryService{
		dao: commodity_dao.NewCommodityCategoryDao(datasource.GetDb()),
	}
}
