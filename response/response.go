package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func WriteJSON(w http.ResponseWriter, statusCode int, resp Response) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(resp)

	if err != nil {
		fmt.Println(err)
	}

}

func WriteError(w http.ResponseWriter, statusCode int, msg string) {
	WriteJSON(w, statusCode, Response{
		Code: statusCode,
		Msg:  msg,
		Data: nil,
	})
}
