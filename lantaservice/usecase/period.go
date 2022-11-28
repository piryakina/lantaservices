package usecase

import (
	"context"
	"lantaservice/entities"
)

type PeriodServer interface {
	AddPeriod(ctx context.Context, p *entities.Period) error
	GetAllPeriods(ctx context.Context) ([]*entities.Period, error)
	GetPeriodNow(ctx context.Context) (*entities.Period, error)
}

type ServicePeriod struct {
	PeriodRepository entities.PeriodRepository
}

func (s *ServicePeriod) AddPeriod(ctx context.Context, p *entities.Period) error {
	err := s.PeriodRepository.AddNewPeriodStorage(ctx, p)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServicePeriod) GetAllPeriods(ctx context.Context) ([]*entities.Period, error) {
	res, err := s.PeriodRepository.GetAllPeriodStorage(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *ServicePeriod) GetPeriodNow(ctx context.Context) (*entities.Period, error) {
	res, err := s.PeriodRepository.GetPeriodNowStorage(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
