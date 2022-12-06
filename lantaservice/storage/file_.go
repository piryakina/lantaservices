package storage

import (
	"context"
	"io"
	"lantaservice/entities"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FileDB struct {
	Id       int64  `db:"id"`
	Filename string `db:"filename"`
	Path     string `db:"path"`
	Date     string `db:"date"`
	Owner    int64  `db:"owner"`
}

func GetFileByOwnerId(ctx context.Context, id int64) {

}
func SaveFile(f multipart.File, header *multipart.FileHeader, fu *entities.File, id int64, status string, idPeriod int64) (*string, error) {
	fullPath := filepath.Join(fu.AbsPath, fu.Folder)
	err := os.MkdirAll(fullPath, 0777)
	if err != nil {
		return nil, err
	}
	fileName := header.Filename
	fileNameRelative := filepath.Join(fu.Folder, fileName)
	fileNameAbs := filepath.Join(fullPath, fileName)
	out, err := os.Create(fileNameAbs)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := out.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	_, err = io.Copy(out, f)
	if err != nil {
		return nil, err
	}
	if strings.Contains(fullPath, "billing") {
		err = SaveBilling(fileName, fileNameRelative, id, status, idPeriod)
		if err != nil {
			return nil, err
		}
		//} else if strings.Contains(fullPath, "test-list") {
		//	err = s.SaveTestList(fileName, fileNameRelative)
		//	if err != nil {
		//		return nil, err
		//	}
	}
	if strings.Contains(fullPath, "invoice") {
		err = SaveInvoice(fileName, fileNameRelative, id, idPeriod)
		if err != nil {
			return nil, err
		}
	}

	return &fileNameRelative, nil
}
func SaveBilling(filename string, path string, id int64, status string, idPeriod int64) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	query := "select id from sp_period where sp=$1 and period=$2"
	row := db.QueryRow(query, id, idPeriod)
	var spPeriodId int64
	if err = row.Scan(&spPeriodId); err != nil {
		return err
	}
	var StatusId int64
	if status != "" {
		query = "select id from docs_status where status_name=$1"
		row = db.QueryRow(query, status)

		if err = row.Scan(&StatusId); err != nil {
			return err
		}
	} else {
		StatusId = 1
	}
	date := time.Now()
	query = "INSERT INTO billing_file (filename, path,status,date, sp_period_id ) VALUES ($1, $2,$3,$4,$5)"
	row = db.QueryRow(query, filename, path, StatusId, date, spPeriodId)
	if err = row.Err(); err != nil {
		return row.Err()
	}
	defer db.Close()
	return nil
}
func SaveInvoice(filename string, path string, id int64, idPeriod int64) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	query := "select id from sp_period where sp=$1 and period=$2"
	row := db.QueryRow(query, id, idPeriod)
	var spPeriodId int64
	if err = row.Scan(&spPeriodId); err != nil {
		return err
	}
	date := time.Now()
	query = "INSERT INTO invoice_file (filename, path,sp_period_id,date) VALUES ($1, $2, $3,$4)"
	row = db.QueryRow(query, filename, path, idPeriod, date)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

//func (s *Storage) SaveTestList(filename string, path string) error {
//	db, err := s.GetDB()
//	if err != nil {
//		return err
//	}
//	query := "INSERT INTO test_list (filename, path) VALUES ($1, $2)"
//	row := db.QueryRow(query, filename, path)
//	if row.Err() != nil {
//		return row.Err()
//	}
//	return nil
//}
