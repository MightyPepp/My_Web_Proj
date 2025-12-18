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

func StartServer(port string) {

	http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/about", aboutHandler)
	// http.HandleFunc("/api/data", apiDataHandler)
	// http.HandleFunc("api/users", usersHandler)

	log.Printf("Сервер запущен на http://localhost:%s", port)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}