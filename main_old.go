package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
    if r.Method != "/home" {
        http.NotFound(w, r)
        return
    }
    w.Write([]byte("Hello this is a go application"))
}

func snippets(w http.ResponseWriter, r *http.Request) {
   
    w.Write([]byte("Sow snippets lists"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        // w.Header().Set("ALLOW", "POST")
        // w.WriteHeader(405)
        // w.Write([]byte("Method not allowed"))
        http.Error(w, "Method not allowed", 405)
        return 
    }
    
    w.Write([]byte("Create snippet here"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))

    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }

    fmt.Fprintf(w,"The id is %d", id)
}


func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/",home)
    mux.HandleFunc("/snippets", snippets)
    mux.HandleFunc("/snippets/create", createSnippet)
    mux.HandleFunc("/snippet", showSnippet)

    log.Println("Starting server on :4000")

    err := http.ListenAndServe(":4000", mux)

    log.Fatal(err)
}