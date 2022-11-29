package webserver

import (
	"encoding/json"
	"lantaservice/entities"
	"lantaservice/usecase"
	"net/http"
	"strconv"
)

func AddUser(w http.ResponseWriter, r *http.Request) { //(s *HttpServer)
	ctx := r.Context()
	c := &entities.User{}
	err := json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	id, err := usecase.SignUpUser(ctx, c)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	JsonResponse(w, StatusResponse{Status: true,
		Detail: strconv.FormatInt(id, 10)}, 200)
	//JsonResponse(w, StatusResponse{Status: true}, 200)
}
