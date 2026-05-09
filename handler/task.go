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
	DeleteTaskByID(id int64) error
	GetTaskDetail(id int64) (*model.TaskDetail, error)
	ListTasksPaginated(userID int64, p model.Pagination) ([]model.Task, int, error)
	UpdateTaskStatus(id int64, status string) error

	// Async test handler
	CompleteTaskAsync(id int64, resultCh chan<- string)
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
	case http.MethodDelete:
		h.DeleteTask(w, r)
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

// GET /tasks?user_id=1&status=&pageSize=&page=
// 改造支持分页和状态过滤
func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")
	status := r.URL.Query().Get("status")

	// 解析 user_id
	var userID int64
	if userIDStr != "" {
		var err error
		userID, err = strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			response.WriteError(w, http.StatusBadRequest, "invalid user_id")
			return
		}
	}

	// 如果没有任何分页参数，保持原来的行为（向后兼容）
	if pageStr == "" && pageSizeStr == "" && status == "" {
		tasks := h.taskService.ListTasksByUserID(userID)
		response.WriteJSON(w, http.StatusOK, response.Response{
			Code: 0,
			Msg:  "ok",
			Data: tasks,
		})
		return
	}

	// 带分页或状态过滤时走新逻辑
	p := model.Pagination{
		Page:     1,
		PageSize: 10,
		Status:   status,
	}
	if pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err != nil || page < 1 {
			response.WriteError(w, http.StatusBadRequest, "invalid page")
			return
		} else {
			p.Page = page
		}
	}
	if pageSizeStr != "" {
		if ps, err := strconv.Atoi(pageSizeStr); err != nil || ps < 1 {
			response.WriteError(w, http.StatusBadRequest, "invalid page_size")
			return
		} else {
			p.PageSize = ps
		}
	}

	tasks, total, err := h.taskService.ListTasksPaginated(userID, p)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Code: 0,
		Msg:  "ok",
		Data: model.PaginatedResult{
			List:     tasks,
			Total:    total,
			Page:     p.Page,
			PageSize: p.PageSize,
		},
	})
}

// DELETE /task { id }
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		response.WriteError(w, http.StatusBadRequest, "id is required")
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "id must be a valid integer")
		return
	}

	if id <= 0 {
		response.WriteError(w, http.StatusBadRequest, "id must be greater than 0")
		return
	}

	if err := h.taskService.DeleteTaskByID(id); err != nil {
		response.WriteError(w, http.StatusInternalServerError, "failed to delete task: "+err.Error())
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Code: 0,
		Msg:  "delete task success",
		Data: map[string]interface{}{
			"deleted_id": id,
		},
	})
}

func (h *TaskHandler) TaskDetail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		response.WriteError(w, http.StatusBadRequest, "id is required")
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "id must be a valid integer")
		return
	}

	taskDetail, err := h.taskService.GetTaskDetail(id)
	if err != nil {
		if err.Error() == "task not found" {
			response.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		response.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Code: 0,
		Msg:  "ok",
		Data: taskDetail,
	})
}

// UpdateTaskStatus 处理 Patch /task 请求，更新任务状态
func (h *TaskHandler) UpdateTaskStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	idStr := r.URL.Query().Get("id")
	status := r.URL.Query().Get("status")

	if idStr == "" || status == "" {
		response.WriteError(w, http.StatusBadRequest, "id and status are required")
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "id must be a valid integer")
		return
	}

	err = h.taskService.UpdateTaskStatus(id, status)
	if err != nil {
		if err.Error() == "task not found" {
			response.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		response.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Code: 0,
		Msg:  "update task status success",
		Data: map[string]interface{}{
			"task_id": id,
			"status":  status,
		},
	})

}

// AsyncTestHandler 仅供开发调试，演示 goroutine + channel 通信
func (h *TaskHandler) AsyncTestHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		response.WriteError(w, http.StatusBadRequest, "id is required")
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "id must be a valid integer")
		return
	}

	resultCh := make(chan string)
	h.taskService.CompleteTaskAsync(id, resultCh)
	result := <-resultCh

	response.WriteJSON(w, http.StatusOK, response.Response{
		Code: 0,
		Msg:  "ok",
		Data: map[string]interface{}{
			"result": result,
		},
	})
}
