package webserver

import (
	"lantaservice/usecase"
	"log"
	"net/http"
	"strconv"
)
//DownloadBilling - скачивание биллинга
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
//DownloadInvoice - скачивание счетов
func DownloadInvoice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	//fmt.Println("id: ", r.Form["id"]) // все получение данных из формы
	fileId, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	filePath, err := usecase.GetInvoicePath(ctx, fileId)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	w.Header().Set("Cache-Control", "no-store")
	//fmt.Println(filePath)
	// JsonResponse(w, StatusResponse{
	// 	Status: true,
	// 	Detail: filePath,
	// }, 200)
	if filePath != "" {
		http.ServeFile(w, r, filePath)
	}
}
//DownloadSLA - скачивание СЛА
func DownloadSLA(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	//fmt.Println("id: ", r.Form["id"]) // все получение данных из формы
	fileId, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	filePath, err := usecase.GetSLAPath(ctx, fileId)
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
//GetImg - получение картинки новости
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
