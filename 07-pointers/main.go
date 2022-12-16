package main

import "fmt"

func main() {
	fmt.Println("welcome to pointer learning")

	// *used for initialization
	// var ptr *int
	// fmt.Println("default value of a pointer is ", ptr)

	number := 23

	// & used for referance
	var numberPointer = &number

	fmt.Println("value of actual pointer is ", numberPointer)
	fmt.Println("value of actual pointer is ", *numberPointer)

	*numberPointer = *numberPointer * 2

	fmt.Println("New value is: ", number)
}