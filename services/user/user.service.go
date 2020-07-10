package user_service

import (
	"errors"
	user_dao "pingduoduo_service/dao/user"
	"pingduoduo_service/datasource"
	user_model "pingduoduo_service/models/user"
)

type UserService interface {
	Login(user_model.User) (*user_model.User, error)
}

type userService struct {
	dao *user_dao.UserDao
}

func (this *userService) Login(param user_model.User) (*user_model.User, error) {
	if len(param.Name) == 0 {
		return nil, errors.New("用户名不能为空")
	}
	if len(param.Password) == 0 {
		return nil, errors.New("密码不能为空")
	}
	return nil, nil
	//var user []user_model.User
	//user = this.dao.Login(param)

}

func NewUserService() UserService {
	return &userService{
		dao: user_dao.NewUserDao(datasource.GetDb()),
	}
}
