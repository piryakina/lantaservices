package storage

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 6543
	//port     = 5432
	user = "postgres"
	//password = "12345678"
	password = "Wt2H1aqF"
	dbname   = "lanta_1"
)

type Storage struct {
	Db *sqlx.DB
}

func GetDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	//db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db, nil

}
