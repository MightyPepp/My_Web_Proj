package server

import (
	"log"
	"net/http"
	"fmt"
)

func StartServer(port string) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/user.html")
	})

	http.HandleFunc("/postform", func(w http.ResponseWriter, r*http.Request) {
		
		username := r.FormValue("username")
		userage := r.FormValue("userage")

		fmt.Fprintf(w, "Имя: %s; Возраст: %s", username, userage)
	})

	log.Printf("Сервер запущен на http://localhost:%s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
	
}