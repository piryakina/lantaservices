package webserver

import (
	"lantaservice/usecase"
	"log"
	"net/http"
	"strconv"
)

func DownloadBilling(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	//fmt.Println("id: ", r.Form["id"]) // все получение данных из формы
	fileId, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	filePath, err := usecase.GetFilePath(ctx, fileId)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	//fmt.Println(filePath)
	// JsonResponse(w, StatusResponse{
	// 	Status: true,
	// 	Detail: filePath,
	// }, 200)
	if filePath != "" {
		http.ServeFile(w, r, filePath)
	}
}

func GetImg(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fileId, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	filePath, err := usecase.GetImgPath(ctx, fileId)
	log.Printf(filePath)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	//fmt.Println(filePath)
	//JsonResponse(w, StatusResponse{
	//	Status: true,
	//	Detail: filePath,
	//}, 200)
	http.ServeFile(w, r, filePath)
}
