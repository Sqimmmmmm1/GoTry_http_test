package repo

import (
	"Gotry_http/model"
	"database/sql"
	"errors"
)

// 定义一个用户仓库对象，它内部持有数据库连接
type UserRepo struct {
	db *sql.DB
}

// 创建一个UserRepo对象 依赖注入，repo本身不会自己创建db链接
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) GetUserByID(id int64) (model.User, error) {
	query := "SELECT id, username, password, age, created_at FROM users WHERE id = ?"

	var user model.User
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Password,
		&user.Age,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, errors.New("user not found")
		}
		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepo) ListUsers() []model.User {
	query := "SELECT id, username, password, age, created_at FROM users ORDER BY id ASC"

	rows, err := r.db.Query(query)
	if err != nil {
		return []model.User{}
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Password,
			&user.Age,
			&user.CreatedAt,
		); err == nil {
			users = append(users, user)
		}
	}
	return users
}

func (r *UserRepo) CreateUser(user model.User) error {
	query := "INSERT INTO users (username, password, age) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, user.Name, user.Password, user.Age)
	return err
}

// func (r *UserRepo) GetUserByName(name string) (model.User, error)
