package webserver

import (
	"encoding/json"
	"lantaservice/entities"
	"net/http"
)

func (s *HttpServer) AddSp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	c := &entities.SP{}
	err := json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	id, err := s.CatalogService.SignUpSP(ctx, c)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	JsonResponse(w, StatusResponse{Status: true,
		Detail: string(id)}, 200)
	JsonResponse(w, StatusResponse{Status: true}, 200)
}
