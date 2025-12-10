package main

import (
	"fmt"
	// "My_Web_Proj/stack"
	"net/http"
)

func main() {
	fmt.Println("---------->><<----------")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Yooooo web")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}