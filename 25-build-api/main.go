package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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

// controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hey wellcome to my world of golang bro.</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all the courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting one specific course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop courses, find specific course by id and return it

	for _, course := range params {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("Course not found with the given id")
	return
}

func createCourse(w http.ResponseWriter, r *http.Request) {

}
