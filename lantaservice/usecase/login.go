package usecase

import (
	"context"
	"lantaservice/entities"
	_ "lantaservice/entities"
	"lantaservice/storage"
)

//type LoginServer interface {
//	Login(ctx context.Context, usr *entities.UserLogin) (int64, error)
//}
//
//type ServiceLogin struct {
//	UserRepository entities.UserRepository
//}

func Login(ctx context.Context, usr *entities.UserLogin) (int64, string, string, error) { //authorization (s *ServiceLogin)
	id, res, role, name, err := storage.LoginUserStorage(ctx, usr.Login)
	if err != nil {
		return 0, "", "", err
	}
	return id, role, name, Compare(res, usr.Password)
}
