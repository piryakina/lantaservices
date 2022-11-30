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
