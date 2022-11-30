package entities

import "time"

type File struct {
	Folder  string
	AbsPath string
}
type DocStatus struct {
	Id         int64  `json:"id,omitempty"`
	StatusName string `json:"status_name"`
}

type Doc struct {
	Id       int64     `json:"id,omitempty"`
	Filename string    `json:"filename,omitempty"`
	Path     string    `json:"path,omitempty"`
	Date     time.Time `json:"date,omitempty"`
	Owner    string    `json:"owner,omitempty"` //владельцем может быть только sp, login sp =owner
	Status   DocStatus `json:"status,omitempty"`
}

//type FileRepository interface {
//	GetFileByOwnerId(ctx context.Context, id int64)
//	SaveFile(f multipart.File, header *multipart.FileHeader, fu *File) (*string, error)
//}
