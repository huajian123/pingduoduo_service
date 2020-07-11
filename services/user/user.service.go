package user_service

import (
	user_dao "pingduoduo_service/dao/user"
	"pingduoduo_service/datasource"
	user_model "pingduoduo_service/models/user"
)

type UserService interface {
	Login(user_model.User) (user_model.User, error)
}

type userService struct {
	dao *user_dao.UserDao
}

func (this *userService) Login(param user_model.User) (user_model.User, error) {
	var user user_model.User
	user = this.dao.Login(param)
	return user, nil
	//var user []user_model.User
	//user = this.dao.Login(param)

}

func NewUserService() UserService {
	return &userService{
		dao: user_dao.NewUserDao(datasource.GetDb()),
	}
}
