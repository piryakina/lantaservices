package main

import (
	"fmt"
	"github.com/rs/cors"
	"lantaservice/webserver"
	"log"
	"net/http"
)

func main() {
	//s := webserver.NewHttpServer()
	Router := webserver.NewRouter()
	//Router.Use(webserver.SessionMiddleware)
	//Router.Use(webserver.MiddlewareCors)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})
	c.Handler(Router)
	http.Handle("/", Router)
	fmt.Println("Сервер запущен")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
