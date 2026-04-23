package handler

import (
	"Gotry_http/model"
	"Gotry_http/response"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type UserService interface {
	GetUserByID(id int64) (model.User, error)
	ListUsers() []model.User
	CreateUser(user model.User) error
}

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetUser start")
	defer fmt.Println("GetUser end")

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

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		if err.Error() == "user not found" {
			response.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		response.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Code: 0,
		Msg:  "ok",
		Data: map[string]interface{}{
			"user": user,
			//"introduce": user.Introduce(),
		},
	})
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ListUsers start")
	defer fmt.Println("ListUsers end")

	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	users := h.userService.ListUsers()

	response.WriteJSON(w, http.StatusOK, response.Response{
		Code: 0,
		Msg:  "ok",
		Data: users,
	})
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req model.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Name == "" {
		response.WriteError(w, http.StatusBadRequest, "name is required")
		return
	}

	if req.Password == "" {
		response.WriteError(w, http.StatusBadRequest, "password is required")
		return
	}

	if req.Age <= 0 {
		response.WriteError(w, http.StatusBadRequest, "age must be greater than 0")
		return
	}

	if err := h.userService.CreateUser(req); err != nil {
		response.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Code: 0,
		Msg:  "create user success",
		Data: map[string]interface{}{
			"username": req.Name,
			"age":      req.Age,
		},
	})
}
