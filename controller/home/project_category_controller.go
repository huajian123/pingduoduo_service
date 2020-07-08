package home_cotrl

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"pingduoduo_service/models"
	home_model "pingduoduo_service/models/home"
	"pingduoduo_service/services/home"
)

// 类声明方法
type ProjectCategoryController struct {
	Ctx     iris.Context
	Service home_service.ProjectCategoryService
}

// 构造函数
func NewProjectCategoryController() *ProjectCategoryController {
	return &ProjectCategoryController{Service: home_service.NewProjectCategoryService()}
}

func (projectCategoryController *ProjectCategoryController) BeforeActivation(a mvc.BeforeActivation) {
	a.Handle("POST", "/queryCategory", "QueryCategory")
	a.Handle("POST", "/addCategory", "AddCategory")
}

// 查询商品类别
func (pc *ProjectCategoryController) QueryCategory() mvc.Result {
	var searchParam models.SearchParam
	err := pc.Ctx.ReadJSON(&searchParam)
	if err != nil {
		pc.Ctx.StatusCode(iris.StatusBadRequest)
		return mvc.Response{
			Object: models.NewResult(nil, 500),
		}
	}
	data := pc.Service.GetProjectCategoryList(searchParam)
	total := pc.Service.GetProjectCategoryListCount()
	return mvc.Response{
		Object: models.NewResult(models.PageInfo{List: data, PageSize: searchParam.PageSize, PageNum: searchParam.PageNum, Total: total}, 0),
	}
}

// 新增商品类别
func (pc *ProjectCategoryController) AddCategory() mvc.Result {
	var addObjModel home_model.ProjectCategory
	err := pc.Ctx.ReadJSON(&addObjModel)
	if err != nil {
		pc.Ctx.StatusCode(iris.StatusBadRequest)
		return mvc.Response{
			Object: models.NewResult(nil, 500),
		}
	}
	pc.Service.AddProjectCategory(addObjModel)
	return mvc.Response{
		Object: models.NewResult(nil, 0),
	}
}
