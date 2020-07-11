package main

import (
	"github.com/kataras/iris/v12"
	"pingduoduo_service/datasource"
	"pingduoduo_service/routes"
)

func main() {
	app := newApp()
	app.Run(iris.Addr(":8081"), iris.WithConfiguration(iris.Configuration{
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "Mon, 02 Jan 2006 15:04:05 GMT",
		Charset:                           "UTF-8",
	}))

}

func newApp() *iris.Application {
	app := iris.New()
	//可以强制应用程序优化以获得尽可能最佳的性能
	app.Configure(iris.WithOptimizations)
	routes.App(app)     //初始化路由
	datasource.InitDb() // 初始化数据库
	return app
}
