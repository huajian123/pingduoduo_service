package list_dao

import (
	"github.com/xormplus/xorm"
	"log"
	list_models "pingduoduo_service/models/list"
)

type SearchListDao struct {
	engine *xorm.Engine
}

func NewSearchListDao(engine *xorm.Engine) *SearchListDao {
	return &SearchListDao{
		engine: engine,
	}
}

func (dao *SearchListDao) GetSearchList() []list_models.SearchList {
	datalist := []list_models.SearchList{}
	err := dao.engine.Find(&datalist)
	if err != nil {
		log.Fatalln(err)
	}
	return datalist
}
