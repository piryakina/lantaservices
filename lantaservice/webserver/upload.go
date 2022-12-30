package webserver

import (
	"encoding/json"
	"lantaservice/entities"
	"lantaservice/usecase"
	"net/http"
	"os"
	"strconv"
	"time"
)

//type uploadedFile struct {
//	ID     int64  `json:"id,omitempty"`
//	status string `json:"status,omitempty"`
//}

// UploadFile - загрузка биллинга на серв
func UploadBilling(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.Body = http.MaxBytesReader(w, r.Body, 150<<20)
	// r.ParseMultipartForm(130<<17)
	f, h, err := r.FormFile("file")
	//defer func() {
	//
	//}()
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	//fmt.Println("id: ", r.Form["id"]) // все получение данных из формы
	strId := r.Form["id"]
	var id int64
	if len(strId) != 0 {
		id, err = strconv.ParseInt(strId[0], 10, 64)
		if err != nil {
			ErrorResponse(w, err)
			return
		}
	}
	status := r.Form["status"]
	var st string
	if len(status) != 0 {
		st = status[0]
	}
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	path, err := os.Getwd()
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	file := entities.File{
		Folder:  "/home/a.piryakina/lanta/lantaservice/lantaservices/lantaservice/upload/billings",
		AbsPath: path,
	}
	date := time.Now()
	//fmt.Println(date)
	res, err := usecase.GetPeriodNow(ctx, date)

	localPath, err := usecase.UploadFile(f, h, &file, id, st, res.Id, 0)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	j, err := json.Marshal(map[string]string{"Url": *localPath})
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	w.Header().Set("ContentBuilder-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(j); err != nil {
		ErrorResponse(w, err)
		return
	}
	if err = f.Close(); err != nil {
		ErrorResponse(w, err)
		return
	}
}

// UploadInvoice UploadFile upload file
func UploadInvoice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.Body = http.MaxBytesReader(w, r.Body, 150<<20)
	f, h, err := r.FormFile("file")

	defer func() {
		if err := f.Close(); err != nil {
			ErrorResponse(w, err)
			return
		}
	}()
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	//c := &entities.User{}
	//err = json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	strId := r.Form["id"]
	var id int64
	if len(strId) != 0 {
		id, err = strconv.ParseInt(strId[0], 10, 64)
		if err != nil {
			ErrorResponse(w, err)
			return
		}
	}
	//status := r.Form["status"]
	//var st string
	//if len(status) != 0 {
	//	st = status[0]
	//}
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	path, err := os.Getwd()
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	date := time.Now()
	//fmt.Println(date)
	res, err := usecase.GetPeriodNow(ctx, date)
	file := entities.File{
		Folder:  "/home/a.piryakina/lanta/lantaservice/lantaservices/lantaservice/upload/invoices",
		AbsPath: path,
	}
	localPath, err := usecase.UploadInvoice(f, h, &file, id, res.Id)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	j, err := json.Marshal(map[string]string{"Url": *localPath})
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	w.Header().Set("ContentBuilder-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(j); err != nil {
		ErrorResponse(w, err)
		return
	}
}
// SetStatusFile - установка статуса файла
func SetStatusFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.URL.Query()
	str := query.Get("statusId")
	str2 := query.Get("fileId")
	stId, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	fileId, err := strconv.ParseInt(str2, 10, 64)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	err = usecase.SetStatus(ctx, stId, fileId)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	JsonResponse(w, StatusResponse{
		Status: true,
		Detail: "status is set",
	}, 200)
}
// GetStatuses - все статусы файлов
func GetStatuses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res, err := usecase.GetStatus(ctx)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	JsonResponse(w, res, 200)
}

//func SetS(w http.ResponseWriter, r *http.Request) {
//	ctx := r.Context()
//	query := r.URL.Query()
//	str := query.Get("id")
//	str2 := query.Get("status")
//	id, err := strconv.ParseInt(str, 10, 64)
//	if err != nil {
//		ErrorResponse(w, err)
//		return
//	}
//
//	err := usecase.SetStatus(ctx, id, str2)
//	if err != nil {
//		ErrorResponse(w, err)
//		return
//	}
//	JsonResponse(w, StatusResponse{
//		Status: true,
//		Detail: "status is set",
//	}, 200)
//}

// UploadAttachments - загрузка вложений в новостям
func UploadAttachments(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	r.Body = http.MaxBytesReader(w, r.Body, 150<<20)
	f, h, err := r.FormFile("file")

	defer func() {
		if err := f.Close(); err != nil {
			ErrorResponse(w, err)
			return
		}
	}()
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	strId := r.Form["id_news"]
	var id int64
	if len(strId) != 0 {
		id, err = strconv.ParseInt(strId[0], 10, 64)
		if err != nil {
			ErrorResponse(w, err)
			return
		}
	}
	//status := r.Form["status"]
	//var st string
	//if len(status) != 0 {
	//	st = status[0]
	//}
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	path, err := os.Getwd()
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	file := entities.File{
		Folder:  "/home/a.piryakina/lanta/lantaservice/lantaservices/lantaservice/upload/attachments",
		AbsPath: path,
	}
	localPath, err := usecase.UploadAttachment(f, h, &file, id)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	j, err := json.Marshal(map[string]string{"Url": *localPath})
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	w.Header().Set("ContentBuilder-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(j); err != nil {
		ErrorResponse(w, err)
		return
	}
}

// UploadFile upload file
func UploadSLA(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.Body = http.MaxBytesReader(w, r.Body, 150<<20)
	// r.ParseMultipartForm(130<<17)
	f, h, err := r.FormFile("file")
	//defer func() {
	//
	//}()
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	//fmt.Println("id: ", r.Form["id"]) // все получение данных из формы
	strId := r.Form["id"]
	strSPId := r.Form["sp"]
	var id, IdSp int64
	if len(strId) != 0 {
		id, err = strconv.ParseInt(strId[0], 10, 64)
		if err != nil {
			ErrorResponse(w, err)
			return
		}
	}
	if len(strSPId) != 0 {
		IdSp, err = strconv.ParseInt(strSPId[0], 10, 64)
		if err != nil {
			ErrorResponse(w, err)
			return
		}
	}
	//status := r.Form["status"]
	//var st string
	//if len(status) != 0 {
	//	st = status[0]
	//}
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	path, err := os.Getwd()
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	file := entities.File{
		Folder:  "/home/a.piryakina/lanta/lantaservice/lantaservices/lantaservice/upload/sla",
		AbsPath: path,
	}
	date := time.Now()
	//fmt.Println(date)
	res, err := usecase.GetPeriodNow(ctx, date)

	localPath, err := usecase.UploadFile(f, h, &file, id, "", res.Id, IdSp)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	j, err := json.Marshal(map[string]string{"Url": *localPath})
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	w.Header().Set("ContentBuilder-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(j); err != nil {
		ErrorResponse(w, err)
		return
	}
	if err = f.Close(); err != nil {
		ErrorResponse(w, err)
		return
	}
}
