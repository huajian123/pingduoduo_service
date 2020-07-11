package commodity_control

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	commodity_service "pingduoduo_service/services/commodity"
)

type CommodityCategoryController struct {
	Ctx     iris.Context
	Service commodity_service.CommodityCategoryService
}

func NewCommodityCategoryController() *CommodityCategoryController {
	return &CommodityCategoryController{Service: commodity_service.NewCommodityCategoryService()}
}

func (projectCategoryController *CommodityCategoryController) BeforeActivation(a mvc.BeforeActivation) {
	a.Handle("POST", "/queryCommodityCategory", "QueryCommodityCategory")
}

func (this *CommodityCategoryController) QueryCommodityCategory() mvc.Result {
	return nil
}
