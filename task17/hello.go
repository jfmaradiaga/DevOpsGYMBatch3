package main

import (
	"fmt"
	"net/http"
	"os"
	"html/template"
)

type HelloRequest struct {
	Name string `json:"name"`
}
//this func helloHandler will handle the form submission, extract the name from the submitted form data and generate the response accordingly
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid form data")
		return
	}

	name := r.FormValue("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Name parameter is missing")
		return
	}

	response := fmt.Sprintf("Hello, %s!", name)
	fmt.Fprintln(w, response)
}
//This indexHandler will server the index.html file, which contains my HTML code/form to read the name, it'll be set to /api/hello 
//indicating that it will be submitted to the helloHandler when we click on the submit button.

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	tmpl := template.Must(template.ParseFiles("index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal server error")
		return
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./cli <host:port>")
		return
	}

	host := os.Args[1]

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/hello", helloHandler)

	err := http.ListenAndServe(host, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

//I have two ways to test this one the first:
// ./cli localhost:8080 Julio where Julio is the variable it reads and prints, the second:
// curl -X GET 'http://localhost:8080/api/hello?name=test' where test is the variable it reads and prints.
//In this updated code, we have added two handlers: indexHandler and helloHandler.
//
//The indexHandler serves the index.html file, which contains a simple HTML form to enter the name. The form action is set to /api/hello, indicating that it will be submitted to the helloHandler when the user clicks the Submit button.
//
//The helloHandler handles the form submission, extracts the name from the submitted form data, and generates the response accordingly.
//
//Make sure to create an index.html file in the same directory as your Go file
