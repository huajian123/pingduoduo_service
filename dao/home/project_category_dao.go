package home_dao

import (
	"github.com/xormplus/xorm"
	"pingduoduo_service/models"
	"pingduoduo_service/models/home"
	"log"
)

type ProjectCategoryDao struct {
	engine *xorm.Engine
}

func NewProjectCategoryDao(engine *xorm.Engine) *ProjectCategoryDao {
	return &ProjectCategoryDao{
		engine: engine,
	}
}

func (p *ProjectCategoryDao) GetProjectCategoryList(param models.SearchParam) []home_model.ProjectCategory {
	dataList := []home_model.ProjectCategory{}
	err := p.engine.Find(&dataList)
	if err != nil {
		log.Fatalln(err)
	}
	return dataList
}

func (p *ProjectCategoryDao) GetProjectCategoryListCount() int {
	count, err := p.engine.Count(new(home_model.ProjectCategory))
	if err != nil {
		log.Fatalln(err)
	}
	return int(count)
}

func (p *ProjectCategoryDao) AddProjectCategory(params home_model.ProjectCategory) {
	count, err := p.engine.Insert(params)
	if err != nil {
		log.Fatalln(err)
	}
	println(count)
	return
}
