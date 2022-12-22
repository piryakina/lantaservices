package usecase

import (
	"context"
	"fmt"
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

func GetFilePath(ctx context.Context, fileId int64) (string, error) {
	doc, err := storage.GetFileInfoById(ctx, fileId)
	if err != nil {
		return "", err
	}
	return doc.Path, nil
}

func GetInvoicePath(ctx context.Context, fileId int64) (string, error) {
	doc, err := storage.GetInvoiceInfoById(ctx, fileId)
	if err != nil {
		return "", err
	}
	return doc.Path, nil
}
func GetSLAPath(ctx context.Context, fileId int64) (string, error) {
	doc, err := storage.GetSLAInfoById(ctx, fileId)
	if err != nil {
		return "", err
	}
	return doc.Path, nil
}
func GetImgPath(ctx context.Context, fileId int64) (string, error) {
	doc, err := storage.GetImgById(ctx, fileId)

	if err != nil {
		return "", err
	}
	if doc != nil {
		fmt.Println(doc.Path)
		return doc.Path, nil

	} else {
		return "/home/a.piryakina/lanta/lantaservice/lantaservices/lantaservice/assets/no-image.svg", nil
	}
}
