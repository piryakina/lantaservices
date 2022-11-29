package entities

type File struct {
	Folder  string
	AbsPath string
}

//type FileRepository interface {
//	GetFileByOwnerId(ctx context.Context, id int64)
//	SaveFile(f multipart.File, header *multipart.FileHeader, fu *File) (*string, error)
//}
