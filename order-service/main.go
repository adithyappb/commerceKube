package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Order Service")
	})
	http.ListenAndServe(":8083", nil)
}

