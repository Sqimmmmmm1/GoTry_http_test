package service

import "Gotry_http/model"

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUserByID(id string) (model.User, error) {
	// TODO: 根据id获取用户信息 先用内存数据返回
	return model.User{}, nil
}

func (s *UserService) GetUserByUsername(username string) (model.User, error) {
	// TODO: 根据用户名获取用户信息 先用内存数据返回
	return model.User{}, nil
}
