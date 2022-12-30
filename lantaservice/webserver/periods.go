package webserver

import (
	"encoding/json"
	"lantaservice/entities"
	"lantaservice/usecase"
	"net/http"
	"time"
)

type EnterPeriod struct {
	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`
	Title    string `json:"title"`
}
//AddNewPeriod - добавление нового периода
func AddNewPeriod(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	c := &entities.Period{}

	enter := &EnterPeriod{}
	err := json.NewDecoder(r.Body).Decode(enter)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	layout := "2006-01-02" //todo yyyy-mm-dd
	c.DateFrom, err = time.Parse(layout, enter.DateFrom)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	c.DateTo, err = time.Parse(layout, enter.DateTo)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	c.Title = enter.Title
	err = usecase.AddPeriod(ctx, c)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	JsonResponse(w, StatusResponse{
		Status: true,
		Detail: "success",
	}, 200)
}
//GetPeriodNow - получение текущего периода
func GetPeriodNow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	date := time.Now()
	//fmt.Println(date)
	res, err := usecase.GetPeriodNow(ctx, date)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	JsonResponse(w, res, 200)
}
//GetAllPeriods - получение всех периодов
func GetAllPeriods(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res, err := usecase.GetAllPeriods(ctx)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	JsonResponse(w, res, 200)
}
