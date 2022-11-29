package webserver

import (
	"encoding/json"
	"lantaservice/entities"
	"lantaservice/usecase"
	"net/http"
	"os"
)

// UploadFile upload file
func UploadBilling(w http.ResponseWriter, r *http.Request) {
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
	path, err := os.Getwd()
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	file := entities.File{
		Folder:  "upload/billings",
		AbsPath: path,
	}
	localPath := usecase.UploadFile(f, h, &file)
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
