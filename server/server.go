package server

import (
	"fmt"
	"net/http"
)

func firstHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "О нас")
}

func StartServer() {

	r := http.NewServeMux()

	// Не работает
	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.HandleFunc("/", firstHandler)

	r.HandleFunc("/about", aboutHandler)

	http.ListenAndServe(":8080", nil)
}
