package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Wt2H1aqF"
	dbname   = "lanta"
)

type Storage struct {
	Db *sqlx.DB
}

func (s *Storage) GetDB() (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//fmt.Println("connected on :5432")
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db, nil

}
