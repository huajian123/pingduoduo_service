package commodity_dao

import (
	"fmt"
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

func (this *CommodityCategoryDao) AddCommodityCategoryList(param commodity_model.CommodityCategory) error {
	commodityCategory := new(commodity_model.CommodityCategory)
	commodityCategory.Name = param.Name
	affected, err := this.engine.Insert(commodityCategory)
	fmt.Println(affected)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (this *CommodityCategoryDao) GetCommodityCategoryList(pageSize int, pageNum int) []commodity_model.CommodityCategory {
	var dataList = make([]commodity_model.CommodityCategory, 0)
	fmt.Println(pageSize)
	fmt.Println(pageNum)
	if pageNum-1 <= 0 {
		pageNum = 1
	}

	err := this.engine.Limit(pageSize*pageNum, (pageNum-1)*pageSize).Find(&dataList)
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
