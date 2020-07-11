package user_control

import (
	"fmt"
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
	a.Handle(iris.MethodPost, "/login", "Login")
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

	if len(user.Password) == 0 || user.Password == "" || user.Name == "" || len(user.Name) == 0 {
		verificatMsg := "用户名或者密码不能为空"
		return mvc.Response{
			Object: models.NewResult(nil, iris.StatusBadRequest, verificatMsg),
		}
	}

	fmt.Println(user)

	response, _ := this.Service.Login(user)
	return mvc.Response{
		Object: models.NewResult(response, 0),
	}
}
