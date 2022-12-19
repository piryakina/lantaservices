package webserver

import (
	"encoding/json"
	"fmt"
	"lantaservice/entities"
	"lantaservice/usecase"
	"net/http"
	"time"
)

func AddNews(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	c := &entities.News{}
	err := json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	date := time.Now().Format("02-01-2006")
	//l := "2001-06-5"
	c.Date = date
	fmt.Println(c.Date)
	id, err := usecase.AddNews(ctx, c)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	JsonResponse(w, StatusId{Id: id}, 200)
}

func GetNews(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var news []*entities.News
	news, err := usecase.GetNews(ctx)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	JsonResponse(w, news, 200)
}
