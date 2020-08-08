package common_controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"io"
	"os"
	"pingduoduo_service/models"
)

type CommonController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

func NewCommodityController() *CommonController {
	return &CommonController{}
}

func (projectCategoryController *CommonController) BeforeActivation(a mvc.BeforeActivation) {
	a.Handle("POST", "/upload", "Upload")
}

func (this *CommonController) Upload() mvc.Result {
	file, info, err := this.Ctx.FormFile("file")
	if err != nil {
		this.Ctx.StatusCode(iris.StatusInternalServerError)
		return mvc.Response{
			Object: models.NewResult(nil, 500),
		}
	}
	defer file.Close()
	fname := info.Filename
	out, err := os.OpenFile("./static/imgs/"+fname,
		os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		this.Ctx.StatusCode(iris.StatusInternalServerError)
		return mvc.Response{
			Object: models.NewResult(nil, 500),
		}
	}
	defer out.Close()
	io.Copy(out, file)
	//ak smAxXNxsdtCRd4ypTFIU3Ucv9mT0NuaFgFGnitCn
	// sk Gsmp9kGk7p4y4au_Go5CeAOjdZdfU2CMIzZ-OtFu
	return mvc.Response{
		Object: "http://localhost:8081/static/imgs/" + fname,
	}
}
