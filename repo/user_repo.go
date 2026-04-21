package repo

import (
	"Gotry_http/model"
	"errors"
)

type UserRepo struct {
	users []model.User
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		users: []model.User{
			{ID: "1", Name: "Tom", Age: 23},
			{ID: "2", Name: "Jack", Age: 25},
			{ID: "3", Name: "Amy", Age: 22},
		},
	}
}

func (r *UserRepo) GetUserByID(id string) (model.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return model.User{}, errors.New("user not found")
}

func (r *UserRepo) ListUsers() []model.User {
	return r.users
}

func (r *UserRepo) CreateUser(user model.User) error {
	for _, u := range r.users {
		if u.ID == user.ID {
			return errors.New("user id already exists")
		}
	}

	r.users = append(r.users, user)
	return nil
}

// func (r *UserRepo) GetUserByName(name string) (model.User, error)
