package entities

import (
	"time"
)

type Period struct {
	Id       int64     `json:"id"`
	DateFrom time.Time `json:"date_from"`
	DateTo   time.Time `json:"date_to"`
	Title    string    `json:"title"`
}

//type PeriodRepository interface {
//	GetPeriodNowStorage(ctx context.Context) (*Period, error)
//	AddNewPeriodStorage(ctx context.Context, p *Period) error
//	GetAllPeriodStorage(ctx context.Context) ([]*Period, error)
//}
