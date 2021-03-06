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

func (this *CommodityCategoryDao) GetCommodityCategoryList(searchParam map[string]interface{}) []commodity_model.CommodityCategory {
	var dataList = make([]commodity_model.CommodityCategory, 0)
	session := this.engine.Where("1=1")
	if searchParam["name"] != "" {
		session = session.And("name = ?", searchParam["name"])
	}
	session.And("pid = ?", searchParam["pid"])

	err := session.Find(&dataList)

	if err != nil {
		log.Fatalln(err)
	}
	return dataList
}

func (this *CommodityCategoryDao) GetCommodityCategoryDetail(id int) commodity_model.CommodityCategory {
	data := commodity_model.CommodityCategory{Id: id}
	_, err := this.engine.Get(&data)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}

func (this *CommodityCategoryDao) UpdateCommodityCategory(param commodity_model.CommodityCategory) commodity_model.CommodityCategory {
	data := commodity_model.CommodityCategory{Name: param.Name}
	_, err := this.engine.ID(param.Id).Update(data)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}

func (this *CommodityCategoryDao) GetSearchListCount(searchParam map[string]interface{}) int {
	session := this.engine.Where("1=1")
	if searchParam["name"] != "" {
		session = session.And("name = ?", searchParam["name"])
	}
	count, err := session.Count(new(commodity_model.CommodityCategory))
	if err != nil {
		log.Fatalln(err)
	}
	return int(count)
}

func (this *CommodityCategoryDao) BatchDelCommdityCategory(ids []int) error {
	obj := new(commodity_model.CommodityCategory)
	_, err := this.engine.In("id", ids).Delete(obj)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (this *CommodityCategoryDao) DelCommodityCategory(id int) error {
	obj := new(commodity_model.CommodityCategory)

	_, err := this.engine.ID(id).Delete(obj)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
