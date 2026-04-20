package main

import "fmt"

// goroutine channel

func sayHello() {
	fmt.Println("Hello from function")
}

func printNum(n int) {
	fmt.Println(n)
}

func producer(ch chan int) {
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	close(ch)
}
