package main

import "fmt"

type Printer interface {
	Print()
}

type Product struct {
	Title string
}

type UserInterface struct {
	name string
}

func (u UserInterface) Print() {
	fmt.Println(u.name)
}

func (p Product) Print() {
	fmt.Println(p.Title)
}
