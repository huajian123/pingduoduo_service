package commodity_dao

import (
	"github.com/xormplus/xorm"
	"log"
	commodity_model "pingduoduo_service/models/commodity"
)

type CommodityCategoryDao struct {
	engine *xorm.Engine
}

func NewCommodityCategoryDao(engine *xorm.Engine) *CommodityCategoryDao {
	return &CommodityCategoryDao{
		engine: engine,
	}
}

func (this *CommodityCategoryDao) GetCommodityCategoryList() []commodity_model.CommodityCategory {
	var dataList = make([]commodity_model.CommodityCategory, 0)
	err := this.engine.Find(&dataList)
	if err != nil {
		log.Fatalln(err)
	}
	return dataList
}

func (this *CommodityCategoryDao) GetSearchListCount() int {
	count, err := this.engine.Count(new(commodity_model.CommodityCategory))
	if err != nil {
		log.Fatalln(err)
	}
	return int(count)
}
