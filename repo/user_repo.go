package repo

import (
	"Gotry_http/model"
	"errors"
)

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) GetUserByID(id string) (model.User, error) {
	// 先用内存模拟
	users := []model.User{
		{ID: "1", Name: "Tom", Age: 23},
		{ID: "2", Name: "Jack", Age: 25},
	}

	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return model.User{}, errors.New("user not found")
}

func (r *UserRepo) ListUsers() []model.User {
	// 先用内存模拟
	return []model.User{
		{ID: "1", Name: "Amy", Age: 23},
		{ID: "2", Name: "Bob", Age: 25},
		{ID: "3", Name: "Carlie", Age: 26},
	}
}

// func (r *UserRepo) GetUserByName(name string) (model.User, error)
