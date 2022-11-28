package storage

import (
	"context"
	"database/sql"
	"lantaservice/entities"
	"time"
)

type PeriodDB struct {
	Id       int64          `json:"id"`
	DateFrom sql.NullTime   `json:"date_from"`
	DateTo   sql.NullTime   `json:"date_to"`
	Title    sql.NullString `json:"title"`
}

func FromPeriodDB(p PeriodDB) *entities.Period {
	var dateFrom time.Time
	if p.DateFrom.Valid {
		dateFrom = p.DateFrom.Time
	}
	var dateTo time.Time
	if p.DateTo.Valid {
		dateTo = p.DateTo.Time
	}
	var title string
	if p.Title.Valid {
		title = p.Title.String
	}
	return &entities.Period{
		Id:       p.Id,
		DateFrom: dateFrom,
		DateTo:   dateTo,
		Title:    title,
	}
}
func (s *Storage) GetPeriodNowStorage(ctx context.Context) (*entities.Period, error) {
	db, err := s.GetDB()
	if err != nil {
		return nil, err
	}
	query := ""
	row := db.QueryRowContext(ctx, query)
	if row.Err() != nil {
		return nil, err
	}
	return nil, nil
}
func (s *Storage) GetAllPeriodStorage(ctx context.Context) ([]*entities.Period, error) {
	db, err := s.GetDB()
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
func (s *Storage) AddNewPeriodStorage(ctx context.Context, p *entities.Period) error {
	db, err := s.GetDB()
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
