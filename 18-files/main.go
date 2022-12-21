package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to files section")
	content := "This needs to go in a file"

	file, err := os.Create("./my-file.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is:", length)
	defer file.Close()

	// reading file method
	readFile("./my-file.txt")
}

// reading file data
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))
}

// checking error
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
