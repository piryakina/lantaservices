package webserver

import (
	"encoding/json"
	"lantaservice/entities"
	"net/http"
)

func (s *HttpServer) AddNewPeriod(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	c := &entities.Period{}
	err := json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	err = s.CatalogService.AddPeriod(ctx, c)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	JsonResponse(w, StatusResponse{
		Status: true,
		Detail: "success",
	}, 200)
}
