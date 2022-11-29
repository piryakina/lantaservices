package storage

import (
	"context"
	"database/sql"
	"lantaservice/entities"
	"time"
)

type NewsDB struct {
	Id    int64          `db:"id"`
	Title sql.NullString `db:"title"`
	Text  sql.NullString `db:"text"`
	Date  sql.NullTime   `db:"date"`
}

func fromNewsDB(p NewsDB) *entities.News {
	var title string
	if p.Title.Valid {
		title = p.Title.String
	}
	var text string
	if p.Text.Valid {
		text = p.Text.String
	}
	var date time.Time
	if p.Date.Valid {
		date = p.Date.Time
	}
	return &entities.News{
		Id:    p.Id,
		Title: title,
		Text:  text,
		Date:  date,
	}
}

func GetNewsStorage(ctx context.Context) ([]*entities.News, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	query := "SELECT title,text,\"date\" from news"
	rows, err := db.QueryContext(ctx, query)
	var news []*entities.News
	for rows.Next() {
		var c NewsDB
		if err = rows.Scan(c.Title, c.Text, c.Date); err != nil {
			return nil, err
		}
		news = append(news, fromNewsDB(c))
	}
	return news, nil
}

func AddNewsStorage(ctx context.Context, p *entities.News) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	query := "INSERT INTO news (title,text,\"date\") VALUES  ($1,$2,$3) returning id"
	title := ToNullString(p.Title)
	text := ToNullString(p.Text)
	date := ToNullTime(p.Date)
	row := db.QueryRowContext(ctx, query, title, text, date)
	var id int64
	if err = row.Scan(&id); err != nil {
		return err
	}
	return nil
}
