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

func (this *UserDao) Login(param user_model.User) user_model.User {
	user := user_model.User{
		Password: param.Password,
		Name:     param.Name,
	}
	has, err := this.engine.Get(&user)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(has)
	return user
}
