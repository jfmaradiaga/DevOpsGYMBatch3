package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./cli <host:port> <name>")
		return
	}

	hostPort := os.Args[1]
	name := os.Args[2]
	url := fmt.Sprintf("http://%s/hello?name=%s", hostPort, name)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making GET request: %s\n", err.Error())
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err.Error())
		return
	}

	fmt.Println(string(body))
}

