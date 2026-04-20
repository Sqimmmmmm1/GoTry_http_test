package repo

import "Gotry_http/model"

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) GetUserByID(id string) (model.User, error) {
	// 先用内存模拟
	return model.User{}, nil
}

func (r *UserRepo) GetUserByUsername(username string) (model.User, error) {
	// 先用内存模拟
	return model.User{}, nil
}
