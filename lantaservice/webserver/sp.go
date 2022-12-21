package webserver

import (
	"encoding/json"
	"fmt"
	"lantaservice/entities"
	"lantaservice/usecase"
	"net/http"
	"time"
)

//func AddSp(w http.ResponseWriter, r *http.Request) {
//	ctx := r.Context()
//	c := &entities.SP{}
//	err := json.NewDecoder(r.Body).Decode(c)
//	if err != nil {
//		ErrorResponse(w, err)
//		return
//	}
//	id, err := usecase.SignUpSP(ctx, c)
//	if err != nil {
//		ErrorResponse(w, err)
//		return
//	}
//	JsonResponse(w, StatusResponse{Status: true,
//		Detail: strconv.FormatInt(id, 10)}, 200)
//}

func GetDataSpPeriodNow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.URL.Query()
	login := query.Get("login")
	date := time.Now()
	res, err := usecase.GetDataSpPeriod(ctx, login, date)
	if res == nil {
		JsonResponse(w, StatusResponse{Status: false}, 200)
	}
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

//	func GetSPNameById(w http.ResponseWriter, r *http.Request) {
//		ctx := r.Context()
//		userSession, err := GetSession(r)
//		name, err := usecase.GetSpNameById(ctx, userSession.ID)
//		if err != nil {
//			JsonResponse(w, StatusResponse{Status: true,
//				Detail: ""}, 200)
//			return
//		}
//		JsonResponse(w, StatusResponse{
//			Status: true,
//			Detail: name,
//		}, 200)
//	}
func GetDataPeriod(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	date := time.Now()
	res, err := usecase.GetPeriodNow(ctx, date)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	fmt.Println("ok down")
	rows, err := usecase.GetDataPeriod(ctx, res.Id)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	JsonResponse(w, rows, 200)
}
func SetCommentFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	c := &entities.CommentFile{}
	err := json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	err = usecase.SetComments(ctx, c.Comment, c.ID)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	JsonResponse(w, StatusResponse{
		Status: true,
		Detail: "comment is set",
	}, 200)
}
