package usecase

import (
	"context"
	"lantaservice/entities"
	_ "lantaservice/entities"
)

type LoginServer interface {
	Login(ctx context.Context, usr *entities.UserLogin) (int64, error)
}

type ServiceLogin struct {
	UserRepository entities.UserRepository
}

func (s *ServiceLogin) Login(ctx context.Context, usr *entities.UserLogin) (int64, error) { //authorization
	id, res, err := s.UserRepository.LoginUserStorage(ctx, usr.Login)
	if err != nil {
		id, res, err = s.UserRepository.LoginSpStorage(ctx, usr.Login)
	}
	if err != nil {
		return 0, err
	}
	return id, Compare(res, usr.Password)
}
