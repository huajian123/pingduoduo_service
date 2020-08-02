package commodity_control

import (
	"fmt"
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
	a.Handle("POST", "/delCommodityCategory", "DelCommodityCategory")
	a.Handle("POST", "/batchDelCommdityCategory", "BatchDelCommdityCategory")
	a.Handle("POST", "/getCommodityCategoryDetail", "GetCommodityCategoryDetail")
	a.Handle("POST", "/updateCommodityCategory", "UpdateCommodityCategory")
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
	dataList := this.Service.GetCommodityCategoryList(searchParam.PageSize, searchParam.PageNum, searchParam.Filters.(map[string]interface{}))
	total := this.Service.GetCommodityCategoryCount(searchParam.Filters.(map[string]interface{}))
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

func (this *CommodityCategoryController) GetCommodityCategoryDetail() mvc.Result {
	var searchParam commodity_model.CommodityCategory
	err := this.Ctx.ReadJSON(&searchParam)
	if err != nil {
		this.Ctx.StatusCode(iris.StatusBadRequest)
		return mvc.Response{
			Object: models.NewResult(nil, 500),
		}
	}
	data := this.Service.GetCommodityCategoryDetail(searchParam.Id)
	return mvc.Response{
		Object: models.NewResult(data, 0),
	}
}

func (this *CommodityCategoryController) UpdateCommodityCategory() mvc.Result {
	var searchParam commodity_model.CommodityCategory
	err := this.Ctx.ReadJSON(&searchParam)
	if err != nil {
		this.Ctx.StatusCode(iris.StatusBadRequest)
		return mvc.Response{
			Object: models.NewResult(nil, 500),
		}
	}
	data := this.Service.UpdateCommodityCategory(searchParam)
	return mvc.Response{
		Object: models.NewResult(data, 0),
	}
}

func (this *CommodityCategoryController) DelCommodityCategory() mvc.Result {
	var delParam commodity_model.CommodityCategory
	err := this.Ctx.ReadJSON(&delParam)
	if err != nil {
		this.Ctx.StatusCode(iris.StatusBadRequest)
		return mvc.Response{
			Object: models.NewResult(nil, 500),
		}
	}
	err2 := this.Service.DelCommodityCategory(delParam.Id)
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

func (this *CommodityCategoryController) BatchDelCommdityCategory() mvc.Result {
	var delParam map[string][]int
	err := this.Ctx.ReadJSON(&delParam)
	if err != nil {
		this.Ctx.StatusCode(iris.StatusBadRequest)
		return mvc.Response{
			Object: models.NewResult(nil, 500),
		}
	}
	err2 := this.Service.BatchDelCommdityCategory(delParam["idArray"])
	fmt.Println(delParam)
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
