package server

// import (
// 	"net/http"
// 	"github.com/gorilla/mux"
// )

// func StartServer() {
// 	r := mux.NewRouter()

// 	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, "index.html")
// 	}).Methods("GET")

// 	http.ListenAndServe(":8080", nil)
// }
