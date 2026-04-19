package main

import (
	"Gotry_http/handler"
	"Gotry_http/model"
	"Gotry_http/response"
	"encoding/json"
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("method:", r.Method)
	//if r.Method != http.MethodGet {
	//	w.WriteHeader(http.StatusMethodNotAllowed)
	//	fmt.Fprintln(w, "method not allowed")
	//	return
	//}
	//
	//fmt.Println("url:", r.URL.Path)

	fmt.Println("收到请求方法了")
	fmt.Println("请求方法", r.Method)
	fmt.Println("请求路径", r.URL.Path)

	fmt.Fprintln(w, "pong")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {

		// w.WriteHeader(http.StatusMethodNotAllowed)
		// fmt.Fprintln(w, "method not allowed")

		response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Code: 0,
		Msg:  "ok",
		Data: map[string]string{
			"message": "hello_day_3",
		},
	})

	//w.Header().Set("Content-Type", "application/json")
	//
	//resp := Response{
	//	Code: 0,
	//	Msg:  "ok",
	//	Data: map[string]string{
	//		"message": "hello",
	//	},
	//}
	//
	//_ = json.NewEncoder(w).Encode(resp)

	//err := json.NewEncoder(w).Encode(resp)
	//if err != nil {
	//	fmt.Println("encode response err:", err)
	//}

}

func userHandler_1(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "method not allowed")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// 1. 从 query 参数中获取id
	id := r.URL.Query().Get("id")

	// 2. 参数校验
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		resp := response.Response{
			Code: 400,
			Msg:  "id is required",
			Data: nil,
		}
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	// 3. 模拟查询用户
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
		{
			w.WriteHeader(http.StatusBadRequest)
			resp := response.Response{
				Code: 404,
				Msg:  "user not found",
				Data: nil,
			}
			_ = json.NewEncoder(w).Encode(resp)
			return
		}
	}

	// 4. 返回成功响应
	resp := response.Response{
		Code: 0,
		Msg:  "ok",
		Data: user,
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func userHandler_2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("userHandler start")
	defer fmt.Println("userHandler end")

	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "id is required")
		return
	}

	var user_1 model.User
	switch id {
	case "1":
		user_1 = model.User{
			ID:   "1",
			Name: "Tom",
			Age:  23,
		}
	case "2":
		user_1 = model.User{
			ID:   "2",
			Name: "Jack",
			Age:  25,
		}
	default:
		{
			response.WriteError(w, http.StatusNotFound, "user not found")
			return
		}
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Code: 0,
		Msg:  "ok",
		Data: map[string]interface{}{
			"user_1":    user_1,
			"introduce": user_1.Introduce(),
		},
	})

}

func usersHandler_1(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "method not allowed")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	users := []model.User{
		{ID: "1", Name: "Amy", Age: 23},
		{ID: "2", Name: "Bob", Age: 25},
		{ID: "3", Name: "Charlie", Age: 26},
	}

	resp := response.Response{
		Code: 0,
		Msg:  "ok",
		Data: users,
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func usersHandler_2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("userHandler start")
	defer fmt.Println("userHandler end")

	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	users2 := []model.User{
		{ID: "1", Name: "Amy", Age: 23},
		{ID: "2", Name: "Bob", Age: 25},
		{ID: "3", Name: "Charlie", Age: 26},
	}

	response.WriteJSON(w, http.StatusOK, response.Response{
		Code: 0,
		Msg:  "ok",
		Data: users2,
	})

}

func main() {
	http.HandleFunc("/ping", handler.PingHandler)
	http.HandleFunc("/hello", handler.HelloHandler)
	http.HandleFunc("/user", handler.UserHandler)

	http.HandleFunc("/users", handler.UsersHandler)

	fmt.Println("server is running at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("start server failed:", err)
	}
}

func day_3_main() {

	user := model.User{
		ID:   "5",
		Name: "Jom",
		Age:  17,
	}

	user.Print()

	user2 := model.User{
		ID:   "6",
		Name: "Jack",
		Age:  18,
	}
	user2.Introduce()

	res, err := divede(10, 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	//res2, err := divede(10, 0)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(res2)

	name, err := getUserByID("1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name)

	err1 := checkAge(1)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println("success")

	var printer Printer

	user3 := UserInterface{name: "Pack"}
	product := Product{Title: "Go Book"}

	printer = user3
	printer.Print()

	printer = product
	printer.Print()

	fmt.Println("start")
	defer fmt.Println("end")
	fmt.Println("running")

	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

}

func test() {
	http.HandleFunc("/ping", pingHandler)

	http.HandleFunc("/hello", helloHandler)

	http.HandleFunc("/user", userHandler_1)

	u := model.User{
		ID:   "1",
		Name: "Qum",
		Age:  21,
	}
	fmt.Println(u)

	fmt.Println(u.ID)
	fmt.Println(u.Name)
	fmt.Println(u.Age)

	u.Name = "Jerry"
	u.Age = 31

	fmt.Println(u)

	names := []string{"Tom", "Jerry", "Alice"}
	for _, name := range names {
		fmt.Println(name)
	}

	names = append(names, "Abc")
	fmt.Println(names)

	users := []model.User{
		{ID: "1", Name: "Tom", Age: 23},
		{ID: "2", Name: "Jerry", Age: 25},
	}
	fmt.Println(users)

	for _, user := range users {
		fmt.Println(user.ID, user.Name, user.Age)
	}

	newUser := model.User{
		ID:   "3",
		Name: "Alice",
		Age:  22,
	}
	users = append(users, newUser)
	fmt.Println(users)

	userInfo := map[string]string{
		"name": "Bill",
		"city": "shanghai",
	}
	fmt.Println(userInfo)

	fmt.Println(userInfo["name"])
	fmt.Println(userInfo["city"])

	fmt.Println("server is running at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("start server failed:", err)
	}
}
