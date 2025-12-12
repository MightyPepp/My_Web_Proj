package server

import (
	"strings"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" { 
		http.ServeFile(w, r, "./static/index.html")
	} else {
		path = strings.TrimPrefix(path, "/")
		http.ServeFile(w, r, "./static/" + path + ".html")
	}
}

func StartServer() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8080", nil)
}