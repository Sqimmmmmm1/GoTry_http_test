package repo

import (
	"Gotry_http/model"
	"database/sql"
)

type TaskRepo struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) CreateTask(task model.Task) error {
	query := "INSERT INTO tasks (user_id, title, status) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, task.UserID, task.Title, task.Status)
	return err
}

func (r *TaskRepo) ListTasksByUserID(userID int64) []model.Task {
	query := "SELECT id, user_id, title, status, created_at, updated_at FROM tasks WHERE user_id = ? ORDER BY id ASC"

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return []model.Task{}
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(
			&task.ID,
			&task.UserID,
			&task.Title,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		); err == nil {
			tasks = append(tasks, task)
		}
	}

	return tasks
}
