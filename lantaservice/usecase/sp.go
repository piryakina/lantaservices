package usecase

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"lantaservice/entities"
	"lantaservice/storage"
	"time"
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

func GetDataSpPeriod(ctx context.Context, login string, date time.Time) (*entities.SpPeriod, error) {
	res, err := storage.GetDataSpPeriodStorage(ctx, login, date)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func AddDataSpPeriod(ctx context.Context, date *entities.SpPeriod) error {
	err := storage.AddDataSpPeriodStorage(ctx, date)
	if err != nil {
		return err
	}
	return nil
}

//func GetSpNameById(ctx context.Context, id int64) (string, error) {
//	name, err := storage.GetSpNameByID(ctx, id)
//	if err != nil {
//		return "", err
//	}
//	return name, nil
//}

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
