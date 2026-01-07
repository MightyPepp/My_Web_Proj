package server

import (
	"log"
	"net/http"
	"strings"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	path = strings.TrimPrefix(path, "/")
	
	http.ServeFile(w, r, "./static/" + path + ".html")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Запрос в корень")
	log.Println("Scheme: ", r.URL.Scheme)
	log.Println("Host: ", r.Host)
	log.Println("Path: ", r.URL.Path)
	log.Println("Query: ", r.URL.Query())
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "./static/index.html")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Запрос в about")
	http.ServeFile(w, r, "./static/about.html")
}

func anotherHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Запрос в another")
	http.ServeFile(w, r, "./static/another.html")
}

func StartServer(port string) {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/another", anotherHandler)

	log.Printf("Сервер запущен на http://localhost:%s", port)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}