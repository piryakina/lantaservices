package webserver

import (
	"encoding/json"
	"lantaservice/entities"
	"lantaservice/usecase"
	"net/http"
	"strconv"
)

func AddSp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	c := &entities.SP{}
	err := json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	id, err := usecase.SignUpSP(ctx, c)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	JsonResponse(w, StatusResponse{Status: true,
		Detail: strconv.FormatInt(id, 10)}, 200)
}
