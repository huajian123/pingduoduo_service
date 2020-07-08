package list_control

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"pingduoduo_service/models"
	list_service "pingduoduo_service/services/list"
)

type SearchListController struct {
	Ctx iris.Context
	Service list_service.SearchListService
}

func NewSearchListController() *SearchListController {
	return &SearchListController{Service: list_service.NewSearchListService()}
}

func (controller *SearchListController) BeforeActivation(a mvc.BeforeActivation) {
	a.Handle("POST", "/list", "GetListData")
}

// 查询商品类别
func (pc *SearchListController) GetListData() mvc.Result {
/*	var searchParam models.SearchParam
	err := pc.Ctx.ReadJSON(&searchParam)
	if err != nil {
		pc.Ctx.StatusCode(iris.StatusBadRequest)
		return mvc.Response{
			Object: models.NewResult(nil, 500),
		}
	}*/
	data := pc.Service.GetSearchList()
	println(&data)
	total := 2
	return mvc.Response{
		Object: models.NewResult(models.PageInfo{List: data, PageSize: 1, PageNum: 1, Total: total},0),
	}
}

