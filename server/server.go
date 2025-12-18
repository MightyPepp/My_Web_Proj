package server

import (
	// "fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Запрос в корень")
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