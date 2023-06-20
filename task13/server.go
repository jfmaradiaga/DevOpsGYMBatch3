package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Create a new HTTP server
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello %s!", r.URL.Query().Get("name"))
    })

    // Listen on port 8080
    http.ListenAndServe(":8080", nil)
}

