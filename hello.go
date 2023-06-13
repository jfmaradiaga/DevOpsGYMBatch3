package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	fmt.Fprintln(w, "Hello, neat AWS EC2 World!")
=======
	fmt.Fprintln(w, "Hello, Neat AWS EC2 World!")
>>>>>>> 54697c47ae3a0e7b57b098426d98f7359184159b
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
