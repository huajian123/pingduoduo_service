package commodity_service

import (
	commodity_dao "pingduoduo_service/dao/commodity"
	"pingduoduo_service/datasource"
	commodity_model "pingduoduo_service/models/commodity"
)

type CommodityCategoryService interface {
	GetCommodityCategoryList() []commodity_model.CommodityCategory
}

type commodityCategoryService struct {
	dao *commodity_dao.CommodityCategoryDao
}

func (this *commodityCategoryService) GetCommodityCategoryList() []commodity_model.CommodityCategory {
	return nil
}

func NewCommodityCategoryService() CommodityCategoryService {
	return &commodityCategoryService{
		dao: commodity_dao.NewCommodityCategoryDao(datasource.GetDb()),
	}
}
