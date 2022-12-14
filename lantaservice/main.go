package main

import (
	"fmt"
	"lantaservice/storage"
	"lantaservice/webserver"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	storage.DBRU = storage.NewStorage()
	Router := webserver.NewRouter()
	Router.Use(webserver.SessionMiddleware)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})

	http.Handle("/", Router)
	fmt.Println("Сервер запущен")
	err := http.ListenAndServe(":8080", c.Handler(Router))
	if err != nil {
		log.Fatal(err)
	}

}
