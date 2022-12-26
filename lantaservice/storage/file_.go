package storage

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"lantaservice/entities"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type StatusDB struct {
	ID         int64  `db:"id"`
	StatusName string `db:"status_name"`
}

func SaveFile(f multipart.File, header *multipart.FileHeader, fu *entities.File, id int64, status string, idPeriod int64, idSp int64) (*string, error) {
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
	if strings.Contains(fullPath, "attachment") {
		err = SaveAttach(fileName, fileNameRelative, id)
		if err != nil {
			return nil, err
		}
	}
	if strings.Contains(fullPath, "sla") {
		err = SaveSLA(fileName, fileNameRelative, id, idPeriod, idSp)
		if err != nil {
			return nil, err
		}
	}
	return &fileNameRelative, nil
}
func SaveBilling(filename string, path string, id int64, status string, idPeriod int64) error {
	db := GetDB()
	query := "select id from sp_period where sp=$1 and period=$2"
	var spPeriodId int64
	err := db.QueryRow(query, id, idPeriod).Scan(&spPeriodId)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no rows with id %d\n", id)
		query = "insert into sp_period (sp,period) values ($1,$2) returning id"
		err = db.QueryRow(query, id, idPeriod).Scan(&spPeriodId)
		if err != nil {
			return err
		}
	case err != nil:
		log.Fatalf("query error: %v\n", err)

	default:
	}
	//if err := row.Scan(&spPeriodId); err != nil {
	//	return err
	//}
	var StatusId int64
	if status != "" {
		query = "select id from docs_status where status_name=$1"
		row := db.QueryRow(query, status)

		if err = row.Scan(&StatusId); err != nil {
			return err
		}
	} else {
		StatusId = 1
	}
	date := time.Now()
	query = "INSERT INTO billing_file (filename, path,status,date, sp_period_id ) VALUES ($1, $2,$3,$4,$5) returning id"
	var fileId int64
	err = db.QueryRow(query, filename, path, StatusId, date, spPeriodId).Scan(&fileId)
	if err != nil {
		return err
	}
	return nil
}
func SaveInvoice(filename string, path string, id int64, idPeriod int64) error {
	db := GetDB()
	query := "select id from sp_period where sp=$1 and period=$2"
	row := db.QueryRow(query, id, idPeriod)
	var spPeriodId int64
	if err := row.Scan(&spPeriodId); err != nil {
		return err
	}
	date := time.Now()
	query = "INSERT INTO invoice_file (filename, path,sp_period_id,date) VALUES ($1, $2, $3,$4) returning id"
	var fileId int64
	err := db.QueryRow(query, filename, path, idPeriod, date).Scan(&fileId)
	if err != nil {
		return err
	}
	return nil
}

func SaveAttach(filename string, path string, id int64) error { //todo attach
	db := GetDB()
	query := "select id from attachment where news_id=$1"
	var attachId int64
	err := db.QueryRow(query, id).Scan(&attachId)
	switch {
	case err == sql.ErrNoRows:
		query = "INSERT INTO attachment (filename, path,news_id) VALUES ($1, $2, $3) returning id"
		var attach int64
		err = db.QueryRow(query, filename, path, id).Scan(&attach)
		if err != nil {
			return err
		}
	case err != nil:
		return err
	default:
		query = "DELETE from attachment where id=$1"
		_, err = db.Exec(query, attachId)
		if err != nil {
			return err
		}
		query = "INSERT INTO attachment (filename, path,news_id) VALUES ($1, $2, $3) returning id"
		var attach int64
		err = db.QueryRow(query, filename, path, id).Scan(&attach)
		if err != nil {
			return err
		}
	}
	return nil
}

func SaveSLA(filename string, path string, id int64, idPeriod int64, idSpPeriod int64) error { //todo attach
	db := GetDB()
	query := "select id from sla_file where sp_period=$1"
	var slaId int64
	err := db.QueryRow(query, idSpPeriod).Scan(&slaId)
	switch {
	case err == sql.ErrNoRows:
		query = "INSERT INTO sla_file (filename, path,usp,sp_period) VALUES ($1, $2, $3, $4) returning id"
		var sla int64
		err = db.QueryRow(query, filename, path, id, idSpPeriod).Scan(&sla)
		if err != nil {
			return err
		}
	case err != nil:
		return err
	default:
		query = "DELETE from sla_file where id=$1"
		_, err = db.Exec(query, slaId)
		if err != nil {
			return err
		}
		query = "INSERT INTO sla_file (filename, path,usp,sp_period) VALUES ($1, $2, $3, $4) returning id"
		var sla int64
		err = db.QueryRow(query, filename, path, id, idSpPeriod).Scan(&sla)
		if err != nil {
			return err
		}
	}
	return nil
}

//	func (s *Storage) SaveTestList(filename string, path string) error {
//		db, err := s.GetDB()
//		if err != nil {
//			return err
//		}
//		query := "INSERT INTO test_list (filename, path) VALUES ($1, $2)"
//		row := db.QueryRow(query, filename, path)
//		if row.Err() != nil {
//			return row.Err()
//		}
//		return nil
//	}

func GetStatusesStorage(ctx context.Context) ([]*entities.DocStatus, error) {
	db := GetDB()
	query := "SELECT id, status_name from docs_status"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var res []*entities.DocStatus
	for rows.Next() {
		var st StatusDB
		if err = rows.Scan(&st.ID, &st.StatusName); err != nil {
			return nil, err
		}
		var item = entities.DocStatus{Id: st.ID, StatusName: st.StatusName}
		res = append(res, &item)
	}
	return res, nil
}

func SetStatusStorage(ctx context.Context, statusId int64, fileId int64) error {
	db := GetDB()
	//query := "select id from user where login=$1"
	//row := db.QueryRowContext(ctx, query, login)
	//var IdSp, spPeriod int64
	//if err = row.Scan(&IdSp); err != nil {
	//	return err
	//}
	//query = "Select id from sp_period where sp=$1 and period=$2"
	//row = db.QueryRowContext(ctx, query, IdSp, idPeriod)
	//if err = row.Scan(&spPeriod); err != nil {
	//	return err
	//}
	query := "update billing_file set status=$1 where id=$2 returning id"
	var idFile int64
	err := db.QueryRowContext(ctx, query, statusId, fileId).Scan(&idFile)
	if err != nil {
		return err
	}
	return nil
}
func SetCommentFile(ctx context.Context, text string, id int64) error {
	db := GetDB()
	query := "update billing_file set comments=$1 where id=$2 returning filename"
	var filename string
	err := db.QueryRowContext(ctx, query, text, id).Scan(&filename)
	if err != nil {
		return err
	}
	return nil
}

func GetFileInfoById(ctx context.Context, id int64) (*entities.BillingFile, error) {
	db := GetDB()
	fmt.Println(id)
	query := "select id, filename, path,date from billing_file where id = $1"
	row := db.QueryRowContext(ctx, query, id)
	var doc BillingFileDB
	if err := row.Scan(&doc.ID, &doc.Filename, &doc.Path, &doc.Date); err != nil {
		return nil, err
	}
	//fmt.Println(doc.ID)
	file := fromFileDB(doc)
	return file, nil
}

func GetInvoiceInfoById(ctx context.Context, id int64) (*entities.InvoiceFile, error) {
	db := GetDB()
	fmt.Println(id)
	query := "select id, filename, path,date from invoice_file where id = $1"
	row := db.QueryRowContext(ctx, query, id)
	var doc InvoiceFileDB
	if err := row.Scan(&doc.ID, &doc.Filename, &doc.Path, &doc.Date); err != nil {
		return nil, err
	}
	//fmt.Println(doc.ID)
	file := fromInvoiceDB(doc)
	return file, nil
}

func GetSLAInfoById(ctx context.Context, id int64) (*entities.SLAFile, error) {
	db := GetDB()
	fmt.Println(id)
	query := "select id, filename, path, usp from sla_file where id = $1"
	row := db.QueryRowContext(ctx, query, id)
	var doc SLAFileDB
	if err := row.Scan(&doc.ID, &doc.Filename, &doc.Path, &doc.USP); err != nil {
		return nil, err
	}
	//fmt.Println(doc.ID)
	file := fromSLADB(doc)
	return file, nil
}

func GetImgById(ctx context.Context, id int64) (*entities.Attach, error) {
	db := GetDB()

	//fmt.Println(id)
	query := "select id, filename, path from attachment where news_id = $1"
	var doc entities.Attach
	err := db.QueryRowContext(ctx, query, id).Scan(&doc.Id, &doc.Filename, &doc.Path)
	log.Printf("kek" + doc.Path)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		log.Fatalf("query error: %v\n", err)
		return nil, err
	default:
		fmt.Printf(doc.Path)
		log.Printf("success")
	}

	//fmt.Println(doc.ID)
	return &doc, nil
}
