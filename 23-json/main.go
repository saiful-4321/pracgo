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

	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
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

func DecodeJson() {
	jsonData := []byte(`
		{
			"coursename": "css",
			"price": 50,
			"website": "w3school.com",
			"tags": ["class", "id"]
		}
	`)

	var mainCourse courses

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("json is valid")
		json.Unmarshal(jsonData, &mainCourse)
		fmt.Printf("%#v\n", mainCourse)
	} else {
		fmt.Println("Json was not valid")
	}

	// data using key value pair

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonData, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is: %v and value is %v and Type is %T\n", k, v, v)
	}
}
