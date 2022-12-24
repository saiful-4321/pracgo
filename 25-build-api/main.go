package main

import "fmt"

// model for course -file
type Course struct {
	CourseId     string  `json:"courseid"`
	CourseName   string  `json:"coursename"`
	CoursePrice  string  `json:"courseprice"`
	CourseAuthor *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware helper - file
func (c *Course) isEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to api calling section")
}
