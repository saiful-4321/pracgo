package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`              //- used to completely vanish any element for user
	Tags     []string `json:"tags,omitempty"` //omitempty to remove element when it is empty
}

func main() {
	fmt.Println("Welcome to json learning world")

	maincourse := []courses{
		{"css", 50, "w3school.com", "abc123", []string{"class", "id"}},
		{"javascript", 500, "w3school.com", "saif4321", []string{"classes", "id", "props", "best"}},
		{"javascript", 500, "w3school.com", "saif4321", nil},
	}

	finalJSON, err := json.MarshalIndent(maincourse, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf(string(finalJSON))

}
