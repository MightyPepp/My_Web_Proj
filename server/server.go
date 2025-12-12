package server

import (
	"net/http"
	"strings"
)

func StartServer() {
	
	r := http.NewServeMux()

	//Хуй пойми в итоге что делает эта строка (либо я тупой)
	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		path = strings.TrimPrefix(path, "/")
		http.ServeFile(w, r, "/static/" + path + ".html")
	})

	http.ListenAndServe(":8080", nil)
}

//Чёт я хуйни наворотил, не работает