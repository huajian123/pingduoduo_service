package commodity_control

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"pingduoduo_service/models"
	commodity_model "pingduoduo_service/models/commodity"
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
	a.Handle("POST", "/addCommodityCategory", "AddCommodityCategory")
}

func (this *CommodityCategoryController) QueryCommodityCategory() mvc.Result {
	var searchParam models.SearchParam
	err := this.Ctx.ReadJSON(&searchParam)
	if err != nil {
		this.Ctx.StatusCode(iris.StatusBadRequest)
		return mvc.Response{
			Object: models.NewResult(nil, 500),
		}
	}
	dataList := this.Service.GetCommodityCategoryList(searchParam.PageSize, searchParam.PageNum)
	total := this.Service.GetCommodityCategoryCount()
	return mvc.Response{
		Object: models.NewResult(models.PageInfo{
			List: dataList, PageNum: searchParam.PageNum, PageSize: searchParam.PageSize, Total: total,
		}, 0),
	}
}

func (this *CommodityCategoryController) AddCommodityCategory() mvc.Result {
	var searchParam commodity_model.CommodityCategory
	err := this.Ctx.ReadJSON(&searchParam)
	if err != nil {
		this.Ctx.StatusCode(iris.StatusBadRequest)
		return mvc.Response{
			Object: models.NewResult(nil, 500),
		}
	}
	err2 := this.Service.AddCommodityCategory(searchParam)
	if err2 != nil {
		this.Ctx.StatusCode(iris.StatusBadRequest)
		return mvc.Response{
			Object: models.NewResult(nil, 500),
		}
	}

	return mvc.Response{
		Object: models.NewResult(nil, 0),
	}
}
