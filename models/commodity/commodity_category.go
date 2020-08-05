package commodity_model

type CommodityCategory struct {
	Id           int    `xorm:"not null pk autoincr INT(11)"  json:"id"`
	Name         string `xorm:"VARCHAR(20)"  json:"name"`
	Pid          int    `xorm:"INT(11)"  json:"pid"`
	CoverUrl     string `xorm:"VARCHAR(255)"  json:"coverUrl"`
	ShowState    int    `xorm:"INT(2)"  json:"showState"`
	SerialNumber int    `xorm:"INT(11)"  json:"serialNumber"`
	Level        int    `xorm:"INT(255)"  json:"level"`
	Sort         int    `xorm:"INT(255)"  json:"sort"`
	Leaf         int    `xorm:"INT(2)"  json:"leaf"`
}
