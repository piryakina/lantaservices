package webserver

import (
	"encoding/json"
	"lantaservice/entities"
	"net/http"
)

func (s *HttpServer) AddNews(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	c := &entities.News{}
	err := json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	err = s.CatalogService.AddNews(ctx, c)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	JsonResponse(w, StatusResponse{
		Status: true,
		Detail: "success",
	}, 200)
}

func (s *HttpServer) GetNews(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var news []*entities.News
	news, err := s.CatalogService.GetNews(ctx)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	JsonResponse(w, news, 200)
}
