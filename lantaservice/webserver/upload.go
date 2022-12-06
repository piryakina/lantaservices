package webserver

import (
	"encoding/json"
	"lantaservice/entities"
	"lantaservice/usecase"
	"net/http"
	"os"
	"fmt"
)

type uploadedFile struct {
	ID       int64  `json:"id,omitempty"`
	status     string `json:"status,omitempty"`
}

// UploadFile upload file
func UploadBilling(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 150<<20)
	// r.ParseMultipartForm(130<<17)
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
	fmt.Println("id: ", r.Form["id"]) // все получение данных из формы
	fmt.Println("status: ", r.Form["status"]) // все получение данных из формы
	c := &entities.User{} 
	err = json.NewDecoder(r.Body).Decode(c.ID) // пытается найти json, которого нет, и поэтому всё ломается
	fmt.Println(err)
	if err != nil {
		fmt.Println("ВЫХОЖУ ТУТ")
		ErrorResponse(w, err)
		return
	}
	fmt.Println("ТУТ2")
	//stat := &entities.DocStatus{} //todo new struct??
	path, err := os.Getwd()
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	file := entities.File{
		Folder:  "upload/billings",
		AbsPath: path,
	}
	localPath := usecase.UploadFile(f, h, &file, c)
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

// UploadInvoice UploadFile upload file
func UploadInvoice(w http.ResponseWriter, r *http.Request) {
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
	c := &entities.User{}
	err = json.NewDecoder(r.Body).Decode(c)
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
		Folder:  "upload/invoice",
		AbsPath: path,
	}
	localPath := usecase.UploadFile(f, h, &file, c)
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
