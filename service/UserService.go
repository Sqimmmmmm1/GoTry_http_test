package service

import "Gotry_http/model"

type UserRepo interface {
	GetUserByID(id string) (model.User, error)
	ListUsers() []model.User
}

type UserService struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUserByID(id string) (model.User, error) {
	// TODO: 根据id获取用户信息 先用内存数据返回
	return s.userRepo.GetUserByID(id)
}

func (s *UserService) ListUsers() []model.User {
	// TODO: 根据用户名获取用户信息 先用内存数据返回
	return s.userRepo.ListUsers()
}
