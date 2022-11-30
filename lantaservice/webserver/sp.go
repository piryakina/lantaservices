package webserver

import (
	"encoding/json"
	"lantaservice/entities"
	"lantaservice/usecase"
	"net/http"
	"strconv"
	"time"
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

func GetDataSpPeriodNow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.URL.Query()
	login := query.Get("login")
	date := time.Now()
	res, err := usecase.GetDataSpPeriod(ctx, login, date)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	JsonResponse(w, res, 200)
}

func AddDataSpPeriodNow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	c := &entities.SpPeriod{}
	err := json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	err = usecase.AddDataSpPeriod(ctx, c)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	JsonResponse(w, StatusResponse{
		Status: true,
		Detail: "success",
	}, 200)
}
