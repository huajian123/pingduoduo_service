package user_dao

import (
	"fmt"
	"github.com/xormplus/xorm"
	"log"
	user_model "pingduoduo_service/models/user"
)

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao{
		engine: engine,
	}
}

func (this *UserDao) Login(param user_model.User) []user_model.User {
	var user []user_model.User
	err := this.engine.Where("name=? and password=?", param.Name, param.Password).Find(&user)
	fmt.Println("22222")
	fmt.Println("%#v\n", user)
	if err != nil {
		log.Fatalln(err)
	}
	return user
}
