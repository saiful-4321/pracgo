package main
import "fmt"

type User struct {
	Name 	string
	Email 	string
	Status 	bool
	Age 	int
}

func main() {
	fmt.Println("Wellcome to struct world")

	Saiful := User{"Saiful", "saiful@dev.com", true, 27}

	fmt.Printf("Saiful details are: %+v\n", Saiful)
	fmt.Printf("Name is: %v\n", Saiful.Name)
}