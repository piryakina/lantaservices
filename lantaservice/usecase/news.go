package usecase

import (
	"context"
	"lantaservice/entities"
)

type NewsServer interface {
	AddNews(ctx context.Context, n *entities.News) error
	GetNews(ctx context.Context) ([]*entities.News, error)
}

type ServiceNews struct {
	NewsRepository entities.NewsRepository
}

func (s *ServiceNews) AddNews(ctx context.Context, n *entities.News) error {
	err := s.NewsRepository.AddNewsStorage(ctx, n)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceNews) GetNews(ctx context.Context) ([]*entities.News, error) {
	news, err := s.NewsRepository.GetNewsStorage(ctx)
	if err != nil {
		return nil, err
	}
	return news, nil
}
