package main

import (
	"Gotry_http/response"
	"encoding/json"
	"fmt"
	"net/http"
)

// 统一 JSON 返回函数
func writeJson(w http.ResponseWriter, statusCode int, resp response.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println(err)
	}
}

// 统一错误返回函数
func writeError(w http.ResponseWriter, statusCode int, msg string) {
	writeJson(w, statusCode, response.Response{
		Code: statusCode,
		Msg:  msg,
		Data: nil,
	})
}
