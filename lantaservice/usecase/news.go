package usecase

import (
	"context"
	"lantaservice/entities"
	"lantaservice/storage"
)

//type NewsServer interface {
//	AddNews(ctx context.Context, n *entities.News) error
//	GetNews(ctx context.Context) ([]*entities.News, error)
//}
//
//type ServiceNews struct {
//	NewsRepository entities.NewsRepository
//}

func AddNews(ctx context.Context, n *entities.News) (int64, error) { //(s *ServiceNews)
	id, err := storage.AddNewsStorage(ctx, n)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetNews(ctx context.Context) ([]*entities.News, error) {

	news, err := storage.GetNewsStorage(ctx)
	if err != nil {
		return nil, err
	}
	return news, nil
}
