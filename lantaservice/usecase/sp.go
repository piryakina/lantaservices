package usecase

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"lantaservice/entities"
	"lantaservice/storage"
)

func SignUpSP(ctx context.Context, usr *entities.SP) (int64, error) { //registration(s *ServiceUser)
	hash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	usr.Password = string(hash)
	id, err := storage.AddSP(ctx, usr)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//func GetSPById(ctx context.Context, usr *entities.SP) (int64, error) { //registration(s *ServiceUser)
//	hash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
//	if err != nil {
//		return 0, err
//	}
//	id, err := storage.SignUpStorage(ctx, usr.Login, string(hash))
//	if err != nil {
//		return 0, err
//	}
//	return id, nil
//}
