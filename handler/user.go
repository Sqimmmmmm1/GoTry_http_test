package handler

import (
	"Gotry_http/model"
	"Gotry_http/response"
	"fmt"
	"net/http"
)

type UserService interface {
	GetUserByID(id string) (model.User, error)
	ListUsers() []model.User
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

	id := r.URL.Query().Get("id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "id is required")
		return
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Code: 0,
		Msg:  "ok",
		Data: map[string]interface{}{
			"user":      user,
			"introduce": user.Introduce(),
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
