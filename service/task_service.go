package service

import (
	"Gotry_http/model"
	"errors"
)

// 定义接口TaskRepo，包含创建任务和按用户ID列表任务的方法
type TaskRepo interface {
	CreateTask(task model.Task) error
	ListTasksByUserID(userID int64) []model.Task
	DeleteTaskByID(id int64) error
	GetTaskByIDWithUser(id int64) (*model.TaskDetail, error)
}

// 定义TaskService结构体，包含TaskRepo接口，把接口作为成员变量，用于依赖注入
type TaskService struct {
	taskRepo TaskRepo
}

// NewTaskService函数用于创建TaskService实例  (taskRepo TaskRepo)表示必须传入一个实现了TaskRepo接口的对象
func NewTaskService(taskRepo TaskRepo) *TaskService {
	return &TaskService{
		taskRepo: taskRepo,
	}
}

// CreateTask方法用于创建任务
func (s *TaskService) CreateTask(task model.Task) error {
	if task.UserID <= 0 {
		return errors.New("user_id must be greater than 0")
	}
	if task.Title == "" {
		return errors.New("title is required")
	}
	if task.Status == "" {
		task.Status = "todo"
	}

	return s.taskRepo.CreateTask(task)
}

// ListTasksByUserID方法用于按用户ID列表任务
func (s *TaskService) ListTasksByUserID(userID int64) []model.Task {
	if userID <= 0 {
		// go中 返回空任务列表
		return []model.Task{}
	}
	return s.taskRepo.ListTasksByUserID(userID)
}

func (s *TaskService) DeleteTaskByID(id int64) error {
	if id <= 0 {
		return errors.New("task id must be greater than 0")
	}
	return s.taskRepo.DeleteTaskByID(id)
}

func (s *TaskService) GetTaskDetail(id int64) (*model.TaskDetail, error) {
	if id <= 0 {
		return nil, errors.New("invalid task id")
	}
	return s.taskRepo.GetTaskByIDWithUser(id) // Changed from s.taskRepo.GetTaskByIDWithUser(id) to s.taskRepo.GetTaskByIDWithUser(id)
}
