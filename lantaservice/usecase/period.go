package usecase

import (
	"context"
	"lantaservice/entities"
	"lantaservice/storage"
	"time"
)

//type PeriodServer interface {
//	AddPeriod(ctx context.Context, p *entities.Period) error
//	GetAllPeriods(ctx context.Context) ([]*entities.Period, error)
//	GetPeriodNow(ctx context.Context) (*entities.Period, error)
//}
//
//type ServicePeriod struct {
//	PeriodRepository entities.PeriodRepository
//}

func AddPeriod(ctx context.Context, p *entities.Period) error { //(s *ServicePeriod)
	err := storage.AddNewPeriodStorage(ctx, p)
	if err != nil {
		return err
	}
	return nil
}

func GetAllPeriods(ctx context.Context) ([]*entities.Period, error) {
	res, err := storage.GetAllPeriodStorage(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func GetPeriodNow(ctx context.Context, date time.Time) (*entities.Period, error) {
	res, err := storage.GetPeriodNowStorage(ctx, date)
	if err != nil {
		return nil, err
	}
	return res, nil
}
