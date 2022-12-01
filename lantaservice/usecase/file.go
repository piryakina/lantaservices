package usecase

import (
	"lantaservice/entities"
	"lantaservice/storage"
	"mime/multipart"
)

//
//type FileServer interface {
//	UploadFile(f multipart.File, h *multipart.FileHeader, fu *entities.File) *string
//}
//type ServiceFile struct {
//	FileRepository entities.FileRepository
//}

// UploadFile upload file
func UploadFile(f multipart.File, h *multipart.FileHeader, fu *entities.File, usr *entities.User) *string { //(s *ServiceFile)
	localPath, err := storage.SaveFile(f, h, fu, usr)
	if err != nil {
		return nil
	}
	return localPath
}

// UploadFile upload file
func UploadInvoice(f multipart.File, h *multipart.FileHeader, fu *entities.File, usr *entities.User) *string { //(s *ServiceFile)
	localPath, err := storage.SaveFile(f, h, fu, usr)
	if err != nil {
		return nil
	}
	return localPath
}
