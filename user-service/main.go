package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "User Service")
	})
	http.ListenAndServe(":8081", nil)
}

