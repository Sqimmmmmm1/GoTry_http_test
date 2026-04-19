package model

import "fmt"

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u User) Introduce() string {
	return fmt.Sprintf("My name is %s, i am %d years old", u.Name, u.Age)
}

// Go语言的规定：方法必须与结构体定义在同一个包中。
func (u User) Print() {
	fmt.Println("ID:", u.ID)
	fmt.Println("name:", u.Name)
	fmt.Println("age:", u.Age)
}
