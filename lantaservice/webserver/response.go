package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//type EquipmentResponse struct {
//	Equipment []*storage.Equipment `json:"equipment"`
//	Status    bool                 `json:"status"`
//}

type StatusResponse struct {
	Status bool   `json:"status"`
	Detail string `json:"detail"`
}
type UserResponse struct {
	Status bool   `json:"status"`
	Id     int64  `json:"id"`
	Role   string `json:"role"`
	Name   string `json:"name"`
}
type StatusId struct {
	Id int64 `json:"id"`
}

func test(w http.ResponseWriter, r *http.Request) {
	JsonResponse(w, "hello", 200)
}

func Index(w http.ResponseWriter, r *http.Request) {
	c, err := GetSession(r)

	if err != nil {
		ErrorResponse(w, err)
		return
	}
	//fmt.Println(c.ID, c.Login)
	JsonResponse(w, StatusId{Id: c.ID}, 200)
}

func JsonResponse(w http.ResponseWriter, js interface{}, status int) {
	b, err := json.Marshal(js)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	w.Header().Set("ContentBuilder-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	_, err = w.Write(b)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
}

// ErrorResponse error handler
func ErrorResponse(w http.ResponseWriter, err error) {
	fmt.Println(err.Error())
	JsonResponse(w, StatusResponse{
		Status: false,
		Detail: err.Error(),
	}, 500)
}
