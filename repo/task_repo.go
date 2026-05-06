package repo

import (
	"Gotry_http/model"
	"database/sql"
)

// 让 TaskRepo 持有数据库连接 （即TaskRepo 就是 专门操作tasks表的数据库助手）
type TaskRepo struct {
	db *sql.DB
}

// 创建一个任务仓库对象，repo自己不创建数据库连接，由main统一创建后注入
func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) CreateTask(task model.Task) error {
	query := "INSERT INTO tasks (user_id, title, status) VALUES (?, ?, ?)"
	// exec 插入操作不需要返回结果集，只关心成不成功
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
