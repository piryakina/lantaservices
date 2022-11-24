package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8080", nil)
}
