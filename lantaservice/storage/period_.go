package storage

import (
	"context"
	"lantaservice/entities"
	"log"
	"time"
)

type PeriodDB struct {
	Id       int64  `json:"id"`
	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`
	Title    string `json:"title"`
}

func FromPeriodDB(p PeriodDB) *entities.Period {
	layout := "2006-01-02 15:04:05Z" //todo yyyy-mm-dd
	var dateFrom time.Time
	dateFrom, err := time.Parse(layout, p.DateFrom)
	if err != nil {
		log.Fatal(err)
	}
	var dateTo time.Time
	dateTo, err = time.Parse(layout, p.DateTo)
	if err != nil {
		log.Fatal(err)
	}
	var title string
	title = p.Title

	return &entities.Period{
		Id:       p.Id,
		DateFrom: dateFrom,
		DateTo:   dateTo,
		Title:    title,
	}
}
func GetPeriodNowStorage(ctx context.Context, date time.Time) (*entities.Period, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	query := "SELECT * FROM period WHERE $1 between date_from and date_to"
	row := db.QueryRowContext(ctx, query, date)
	var period *entities.Period
	var temp PeriodDB
	if err = row.Scan(&temp.Id, &temp.DateFrom, &temp.DateTo, &temp.Title); err != nil {
		return nil, err
	}
	period = FromPeriodDB(temp)
	return period, nil
}
func GetAllPeriodStorage(ctx context.Context) ([]*entities.Period, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	query := "SELECT * from period"
	rows, err := db.QueryContext(ctx, query)
	var prds []*entities.Period
	for rows.Next() {
		var c PeriodDB
		if err = rows.Scan(&c.Id, &c.DateFrom, &c.DateTo, &c.Title); err != nil {
			return nil, err
		}
		prds = append(prds, FromPeriodDB(c))
	}
	defer rows.Close()
	return prds, nil
}
func AddNewPeriodStorage(ctx context.Context, p *entities.Period) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	query := "INSERT INTO period (date_from,date_to,title) VALUES ($1,$2,$3) returning id"
	row := db.QueryRowContext(ctx, query, p.DateFrom, p.DateTo, p.Title)
	var id int64
	if err = row.Scan(&id); err != nil {
		return err
	}
	return nil
}
