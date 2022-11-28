package storage

import "context"

type FileDB struct {
	Id       int64  `db:"id"`
	Filename string `db:"filename"`
	Path     string `db:"path"`
	Date     string `db:"date"`
	Owner    int64  `db:"owner"`
}

func (s *Storage) GetFileByOwnerId(ctx context.Context, id int64) {

}
