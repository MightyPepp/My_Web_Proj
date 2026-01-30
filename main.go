package main

import (
	"net/http"
)

func coreHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", coreHandler)

	http.ListenAndServe(":8080", mux)
}