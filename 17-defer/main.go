package main

import "fmt"

func main() {
	defer fmt.Println("defer exaples1")
	defer fmt.Println("defer exaples2")
	fmt.Println("Welcome to defer section")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("Number is: %v\n", i)
	}
}
