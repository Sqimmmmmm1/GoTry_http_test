package handler

import (
	"Gotry_http/model"
	"Gotry_http/response"
	"encoding/json"
	"net/http"
	"strconv"
)

type TaskService interface {
	CreateTask(task model.Task) error
	ListTasksByUserID(userID int64) []model.Task
}

type TaskHandler struct {
	taskService TaskService
}

func NewTaskHandler(taskService TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

// 总入口：按请求方法分发
func (h *TaskHandler) Tasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateTask(w, r)
	case http.MethodGet:
		h.ListTasks(w, r)
	default:
		response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

// POST /tasks
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req model.Task
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.UserID <= 0 {
		response.WriteError(w, http.StatusBadRequest, "user_id must be greater than 0")
		return
	}

	if req.Title == "" {
		response.WriteError(w, http.StatusBadRequest, "title is required")
		return
	}

	// 为了让返回值更直观，先在 handler 里补默认值
	// 你 service 里也可以保留同样逻辑，当前阶段重复一点没问题
	if req.Status == "" {
		req.Status = "todo"
	}

	if err := h.taskService.CreateTask(req); err != nil {
		response.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Code: 0,
		Msg:  "create task success",
		Data: map[string]interface{}{
			"user_id": req.UserID,
			"title":   req.Title,
			"status":  req.Status,
		},
	})
}

// GET /tasks?user_id=1
func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		response.WriteError(w, http.StatusBadRequest, "user_id is required")
		return
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "user_id must be a valid integer")
		return
	}

	tasks := h.taskService.ListTasksByUserID(userID)

	response.WriteJSON(w, http.StatusOK, response.Response{
		Code: 0,
		Msg:  "ok",
		Data: tasks,
	})
}
