package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

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

	response := HelloResponse{Message: fmt.Sprintf("Hello, %s!", name)}
	jsonData, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error generating response")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./cli <host:port>")
		return
	}

	host := os.Args[1]

	http.HandleFunc("/api/hello", helloHandler)

	err := http.ListenAndServe(host, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

