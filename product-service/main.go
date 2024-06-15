package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Product Service")
	})
	http.ListenAndServe(":8082", nil)
}

