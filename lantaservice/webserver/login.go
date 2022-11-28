package webserver

import (
	"encoding/json"
	"lantaservice/entities"
	"net/http"
)

func (s *HttpServer) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	c := &entities.UserLogin{}
	err := json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	var id int64
	id, err = s.CatalogService.Login(ctx, c)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	user := SessionContext{
		ID:    id,
		Login: c.Login,
	}
	err = SetSession(w, user)
	if err != nil {
		ErrorResponse(w, err)
		return
	}

	JsonResponse(w, StatusResponse{
		Status: true, Detail: "session is active",
	}, 200)
}
func Logout(w http.ResponseWriter, r *http.Request) {
	DeleteSession(w)
	JsonResponse(w, StatusResponse{Status: true}, 200)
}
