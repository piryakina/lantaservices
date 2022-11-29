package webserver

import (
	"encoding/json"
	"lantaservice/entities"
	"lantaservice/usecase"
	"net/http"
)

func AddNews(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	c := &entities.News{}
	err := json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	err = usecase.AddNews(ctx, c)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	JsonResponse(w, StatusResponse{
		Status: true,
		Detail: "success",
	}, 200)
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
