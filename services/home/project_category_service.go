package home_service

import (
	home_dao "pingduoduo_service/dao/home"
	"pingduoduo_service/datasource"
	"pingduoduo_service/models"
	"pingduoduo_service/models/home"
)

// 类
type projectCategoryService struct {
	dao *home_dao.ProjectCategoryDao
}

// 构造函数
func NewProjectCategoryService() ProjectCategoryService {
	return &projectCategoryService{
		dao: home_dao.NewProjectCategoryDao(datasource.GetDb()),
	}
}

// 接口
type ProjectCategoryService interface {
	GetProjectCategoryList(models.SearchParam) []home_model.ProjectCategory
	GetProjectCategoryListCount() int
	AddProjectCategory(category home_model.ProjectCategory)
}

func (p projectCategoryService) GetProjectCategoryList(params models.SearchParam) []home_model.ProjectCategory {
	return p.dao.GetProjectCategoryList(params)
}

func (p projectCategoryService) GetProjectCategoryListCount() int {
	return p.dao.GetProjectCategoryListCount()
}

func (p projectCategoryService) AddProjectCategory(params home_model.ProjectCategory) {
	p.dao.AddProjectCategory(params)
	return
}
