package service

import (
	"Gotry_http/model"
	"errors"
)

type TaskRepo interface {
	CreateTask(task model.Task) error
	ListTasksByUserID(userID int64) []model.Task
}

type TaskService struct {
	taskRepo TaskRepo
}

func NewTaskService(taskRepo TaskRepo) *TaskService {
	return &TaskService{
		taskRepo: taskRepo,
	}
}

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

func (s *TaskService) ListTasksByUserID(userID int64) []model.Task {
	if userID <= 0 {
		return []model.Task{}
	}
	return s.taskRepo.ListTasksByUserID(userID)
}
