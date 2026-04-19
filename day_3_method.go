package main

import "fmt"

func (u User) Print() {
	fmt.Println("ID:", u.ID)
	fmt.Println("name:", u.Name)
	fmt.Println("age:", u.Age)
}

func (u User) Introduce() string {
	return fmt.Sprintf("My name is %s, i am %d years old", u.Name, u.Age)
}
