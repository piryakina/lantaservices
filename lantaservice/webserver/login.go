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
	var id int64
	id, err = usecase.Login(ctx, c)
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
