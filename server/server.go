package server

import (
	"log"
	"net/http"
	"fmt"

	"math/rand"
	"strconv"
)

func StartServer(port string) {


//----------------------------------------------------------------------------
	http.HandleFunc("/casinoPostform", func(w http.ResponseWriter, r *http.Request) {

		usernumber := r.FormValue("usernumber")

		number, err := strconv.Atoi(usernumber)

		if err != nil {
			fmt.Fprintf(w, "Ошибка блять %s", err)
		} else {
			if 0 <= number && number <= 10 {
				casinoNumber := rand.Intn(11)
				if number == casinoNumber {
					w.Header().Set("Content-Type", "text/html; charset=utf-8")
					html := `
					<!DOCTYPE html>
					<html>
						<head>
							<meta charset = "utf-8">
							<tittle>Win page</tittle>
						</head>
						<body>
							<div>
								<h1>You win!</h1>
								<a href="http://localhost:8080/casino">Go back and try again</a>
							</div>
						</body>
					</html>
					`
					w.Write([]byte(html))
				} else {
					fmt.Fprint(w, "Вы проиграли, казино загадало число ", casinoNumber)
				}
			} else {
				fmt.Fprint(w, "Число должно быть от 0 до 10!")
			}
		}

	})

	http.HandleFunc("/casino", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/casino.html")
	})
//----------------------------------------------------------------------------


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