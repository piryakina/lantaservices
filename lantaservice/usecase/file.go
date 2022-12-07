package usecase

import (
	"context"
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
func UploadInvoice(f multipart.File, h *multipart.FileHeader, fu *entities.File, id int64, idPeriod int64) (*string, error) { //(s *ServiceFile)
	localPath, err := storage.SaveFile(f, h, fu, id, "", idPeriod)
	if err != nil {
		return nil, err
	}
	return localPath, err
}
func UploadAttachment(f multipart.File, h *multipart.FileHeader, fu *entities.File, id int64) (*string, error) { //(s *ServiceFile)
	localPath, err := storage.SaveFile(f, h, fu, id, "", 0)
	if err != nil {
		return nil, err
	}
	return localPath, err
}

func SetStatus(ctx context.Context, idSt int64, idFile int64) error {
	err := storage.SetStatusStorage(ctx, idSt, idFile)
	return err
}
func GetStatus(ctx context.Context) ([]*entities.DocStatus, error) {
	res, err := storage.GetStatusesStorage(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func SetComments(ctx context.Context, text string, id int64) error {
	err := storage.SetCommentFile(ctx, text, id)
	return err
}
