package main

import (
	"fmt"
	"github.com/rs/cors"
	"lantaservice/webserver"
	"log"
	"net/http"
)

func main() {

	Router := webserver.NewRouter()
	Router.Use(webserver.SessionMiddleware)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200", "https://dbc3-94-25-185-137.eu.ngrok.io"},
		AllowCredentials: true,
	})

	http.Handle("/", Router)
	fmt.Println("Сервер запущен")
	err := http.ListenAndServe(":8080", c.Handler(Router))
	if err != nil {
		log.Fatal(err)
	}

}
