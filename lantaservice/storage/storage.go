package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"

	_ "github.com/lib/pq"
)

var DBRU *sqlx.DB

const (
	port = 6543
	//port = 5432
	user = "postgres"
	//password = "12345678"
	password = "Wt2H1aqF"
	//password = "winx2011"
	dbname = "lanta_1"
)

type Storage struct {
	Db *sqlx.DB
}

// NewStorage constructor for storage
func NewStorage() *sqlx.DB {
	dbRu, err := InitDB()
	if err != nil {
		return nil
	}
	return dbRu
}
func InitDB() (*sqlx.DB, error) {
	host := os.Getenv("host")
	fmt.Println(host)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sqlx.Connect("postgres", psqlInfo)
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
func GetDB() *sqlx.DB {
	return DBRU
}
