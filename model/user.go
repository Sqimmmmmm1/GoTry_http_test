package model

import (
	"fmt"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"username"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

//func (u User) Introduce() string {
//	return fmt.Sprintf("My name is %s, i am %d years old", u.Name, u.Age)
//}

// Go语言的规定：方法必须与结构体定义在同一个包中。
func (u User) Print() {
	fmt.Println("ID:", u.ID)
	fmt.Println("name:", u.Name)
	fmt.Println("age:", u.Age)
}
