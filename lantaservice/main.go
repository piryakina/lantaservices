package main

import (
	"github.com/rs/cors"
	"lantaservice/webserver"
	"log"
	"net/http"
)

func main() {
	Router := webserver.NewRouter()
	Router.Use(webserver.SessionMiddleware)
	//Router.Use(webserver.MiddlewareCors)
	http.Handle("/", Router)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})
	err := http.ListenAndServe(":8080", c.Handler(Router))
	if err != nil {
		log.Fatal(err)
	}
}
