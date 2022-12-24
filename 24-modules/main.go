package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello from mod in golang")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func greeter() {
	fmt.Println("Hey there mod uesers.")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello from the other side</h1>"))
}
