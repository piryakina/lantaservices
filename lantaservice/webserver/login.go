package webserver

import (
	"encoding/json"
	"lantaservice/entities"
	"lantaservice/usecase"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) { //(s *HttpServer)
	ctx := r.Context()
	c := &entities.UserLogin{}
	err := json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	id, role, name, err := usecase.Login(ctx, c)
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

	JsonResponse(w, UserResponse{
		Status: true,
		Id:     id,
		Role:   role,
		Name:   name,
	}, 200)
}
func Logout(w http.ResponseWriter, r *http.Request) {
	DeleteSession(w)
	JsonResponse(w, StatusResponse{Status: true}, 200)
}
