package handler

import (
	"Gotry_http/model"
	"Gotry_http/response"
	"fmt"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UserHandler start")
	defer fmt.Println("UserHandler end")

	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "id is required")
		return
	}

	var user model.User
	switch id {
	case "1":
		user = model.User{
			ID:   "1",
			Name: "Tom",
			Age:  23,
		}
	case "2":
		user = model.User{
			ID:   "2",
			Name: "Jack",
			Age:  25,
		}
	default:
		response.WriteError(w, http.StatusNotFound, "user not found")
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

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UsersHandler start")
	defer fmt.Println("UsersHandler end")

	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	users := []model.User{
		{ID: "1", Name: "Amy", Age: 23},
		{ID: "2", Name: "Bob", Age: 25},
		{ID: "3", Name: "Carlie", Age: 26},
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Code: 0,
		Msg:  "ok",
		Data: users,
	})

}
