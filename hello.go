package main

import (
        "fmt"
        "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, Neat AWS EC2 World!")
}

func main() {
        http.HandleFunc("/", helloHandler)
        http.ListenAndServe(":8080", nil)
}
