package service

import (
	"Gotry_http/model"
	"errors"
)

type UserRepo interface {
	GetUserByID(id int64) (model.User, error)
	ListUsers() []model.User
	CreateUser(user model.User) error
}

type UserService struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUserByID(id int64) (model.User, error) {
	// TODO: 根据id获取用户信息 先用内存数据返回
	if id <= 0 {
		return model.User{}, errors.New("invalid id")
	}
	return s.userRepo.GetUserByID(id)
}

func (s *UserService) ListUsers() []model.User {
	// TODO: 根据用户名获取用户信息 先用内存数据返回
	return s.userRepo.ListUsers()
}

func (s *UserService) CreateUser(user model.User) error {
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	if user.Age <= 0 {
		return errors.New("age must be greater than 0")
	}

	return s.userRepo.CreateUser(user)
}
