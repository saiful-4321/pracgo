package main
import "fmt"

func main() {
	fmt.Println("wellcome to function section")

	greeter()

	result := adder(3, 5)

	proRes, myMessage := proAdder(9,8,7)

	fmt.Println("Result is:", result)
	fmt.Println(myMessage, proRes)
}

func greeter() {
	fmt.Println("Hello from the other side")
}

func adder(first int, secound int) int {
	return first + secound
}

func proAdder(values ...int) (int,string) {
	total := 0

	for _, value := range values {
		total += value
	}

	return total, "Pro result should be: "
}