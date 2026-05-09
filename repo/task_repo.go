package repo

import (
	"Gotry_http/model"
	"database/sql"
	"errors"
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

func (r *TaskRepo) DeleteTaskByID(id int64) error {
	query := "DELETE FROM tasks WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}

// GetTaskByIDWithUser 通过任务ID获取任务，同时LEFT JOIN 查询任务及所属用户
func (r *TaskRepo) GetTaskByIDWithUser(id int64) (*model.TaskDetail, error) {
	query := "SELECT t.id, t.user_id, t.title, t.status, t.created_at, t.updated_at, u.id, u.username, u.age, u.created_at FROM tasks t LEFT JOIN users u ON t.user_id = u.id WHERE t.id = ?"

	var task model.TaskDetail
	var user model.User

	err := r.db.QueryRow(query, id).Scan(
		&task.ID,
		&task.UserID,
		&task.Title,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
		&user.ID,
		&user.Name,
		&user.Age,
		&user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}
	if user.ID > 0 {
		task.User = &user
	}
	return &task, nil
}

// ListTaskWithFiilter 支持user_id、状态过滤、分页
func (r *TaskRepo) ListTasksWithFilter(userID int64, status string, offset, limit int) ([]model.Task, error) {
	query := "SELECT id, user_id, title, status, created_at, updated_at FROM tasks WHERE 1=1"
	args := []interface{}{}

	if userID > 0 {
		query += " AND user_id = ?"
		args = append(args, userID)
	}

	if status != "" {
		query += " AND status = ?"
		args = append(args, status)
	}

	query += " ORDER BY id ASC LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		if err := rows.Scan(
			&t.ID,
			&t.UserID,
			&t.Title,
			&t.Status,
			&t.CreatedAt,
			&t.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

// CountTasks 统计符合条件的任务总数
func (r *TaskRepo) CountTasks(userID int64, status string) (int, error) {
	query := "SELECT COUNT(*) FROM tasks WHERE 1=1"
	args := []interface{}{}

	if userID > 0 {
		query += " AND user_id = ?"
		args = append(args, userID)
	}

	if status != "" {
		query += " AND status = ?"
		args = append(args, status)
	}

	var count int
	err := r.db.QueryRow(query, args...).Scan(&count)
	return count, err
}

// UpdateTaskStatus 更新任务状态
func (r *TaskRepo) UpdateTaskStatus(id int64, status string) error {
	query := "UPDATE tasks SET status = ? WHERE id = ?"
	res, err := r.db.Exec(query, status, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("task not found")
	}
	return nil
}
