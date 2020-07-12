package commodity_model

type CommodityCategory struct {
	Id   int    `xorm:"not null pk INT(11)"  json:"id"`
	Name string `xorm:"VARCHAR(20)"  json:"name"`
}
