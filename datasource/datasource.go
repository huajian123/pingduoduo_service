package datasource

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

var engine *xorm.Engine

func GetDb() *xorm.Engine {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:abc15150568124@/hjdb?charset=utf8")
	if err != nil {
		println(err.Error())
		return engine
	}
	engine.ShowSQL(true)
	return engine
}

func InitDb() {

}
