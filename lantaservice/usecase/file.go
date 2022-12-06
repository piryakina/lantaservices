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
func UploadFile(f multipart.File, h *multipart.FileHeader, fu *entities.File, id int64, status string, idPeriod int64) (*string, error) { //(s *ServiceFile)
	localPath, err := storage.SaveFile(f, h, fu, id, status, idPeriod)
	if err != nil {
		return nil, err
	}
	return localPath, err
}

// UploadFile upload file
func UploadInvoice(f multipart.File, h *multipart.FileHeader, fu *entities.File, id int64, status string, idPeriod int64) (*string, error) { //(s *ServiceFile)
	localPath, err := storage.SaveFile(f, h, fu, id, status, idPeriod)
	if err != nil {
		return nil, err
	}
	return localPath, err
}
