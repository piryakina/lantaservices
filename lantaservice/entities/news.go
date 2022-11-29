package entities

import (
	"time"
)

type News struct {
	Id    int64     `json:"id"`
	Title string    `json:"title"`
	Text  string    `json:"text"`
	Date  time.Time `json:"date"`
}

//type NewsRepository interface {
//	GetNewsStorage(ctx context.Context) ([]*News, error)
//	AddNewsStorage(ctx context.Context, p *News) error
//}
