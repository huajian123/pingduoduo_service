package user_control

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"pingduoduo_service/models"
	user_model "pingduoduo_service/models/user"
	user_service "pingduoduo_service/services/user"
)

type UserController struct {
	Ctx     iris.Context
	Service user_service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		Service: user_service.NewUserService(),
	}
}

func (controller *UserController) BeforeActivation(a mvc.BeforeActivation) {
	a.Handle("POST", "/login", "Login")
}

func (this *UserController) Login() mvc.Result {
	var user user_model.User
	err := this.Ctx.ReadJSON(&user)
	if err != nil {
		this.Ctx.StatusCode(iris.StatusBadRequest)
		return mvc.Response{
			Object: models.NewResult(nil, 500),
		}
	}
	response, _ := this.Service.Login(user)
	return mvc.Response{
		Object: models.NewResult(response, 0),
	}
}
