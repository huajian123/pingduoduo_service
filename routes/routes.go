package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	home_cotrl "pingduoduo_service/controller/home"
	list_control "pingduoduo_service/controller/list"
	user_control "pingduoduo_service/controller/user"
)

func App(app *iris.Application) {

	mvc.Configure(app.Party("/home"), func(context *mvc.Application) {
		context.Handle(home_cotrl.NewProjectCategoryController())
	})

	mvc.Configure(app.Party("/ceshilist"), func(context *mvc.Application) {
		context.Handle(list_control.NewSearchListController())
	})

	mvc.Configure(app.Party("/user"), func(context *mvc.Application) {
		context.Handle(user_control.NewUserController())
	})
}
