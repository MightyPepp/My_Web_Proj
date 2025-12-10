package main

import (
	"fmt"
	// "My_Web_Proj/stack"
	"net/http"
)

func main() {
	fmt.Println("---------->><<----------")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Yooooo web!")
		// http.ServeFile(w, r, "index.html") //Нужно разобраться, что делать если .html в какой-то папке лежит
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}