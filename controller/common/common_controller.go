package common_controller

import (
	"context"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"io"
	"pingduoduo_service/models"
	"strconv"
	"time"
)

const (
	accessKey = "smAxXNxsdtCRd4ypTFIU3Ucv9mT0NuaFgFGnitCn"
	secretKey = "Gsmp9kGk7p4y4au_Go5CeAOjdZdfU2CMIzZ-OtFu"
	bucket    = "pingduodspace"
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

func config() storage.Config {
	cfg := storage.Config{}

	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	return cfg
}

func Upload(localFile io.Reader, size int64, filename string) (string, error) {

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}

	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cig := config()

	formUploader := storage.NewFormUploader(&cig)

	ret := storage.PutRet{}

	putExtra := storage.PutExtra{}

	err := formUploader.Put(context.Background(), &ret, upToken, filename, localFile, size, &putExtra)

	if err != nil {
		return "", err
	}

	return ret.Key, nil
}

func (this *CommonController) Upload() mvc.Result {
	f, h, err := this.Ctx.FormFile("file")
	if err != nil {
		this.Ctx.StatusCode(iris.StatusInternalServerError)
		return mvc.Response{
			Object: models.NewResult(nil, 500),
		}
	}
	defer f.Close()
	file, _ := h.Open()
	unix := time.Now().Unix()
	timeByte := []byte(strconv.Itoa(int(unix)))
	filename := string(timeByte)

	if filename, err = Upload(file, h.Size, filename); err != nil { // 通过h.size 即可获得文件大小
		fmt.Println("成功")
		return nil
	} else {
		fmt.Println("失败")
		return nil
	}

	/*out, err := os.OpenFile("./static/imgs/"+fname,
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
	}*/
}
