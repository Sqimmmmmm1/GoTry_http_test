package handler

import (
	"Gotry_http/response"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Code: 0,
		Msg:  "ok",
		Data: map[string]string{
			"message": "hello",
		},
	})
}
