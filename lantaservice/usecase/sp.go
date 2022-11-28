package usecase

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"lantaservice/entities"
)

func (s *ServiceUser) SignUpSP(ctx context.Context, usr *entities.SP) (int64, error) { //registration
	hash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	id, err := s.UserRepository.SignUpStorage(ctx, usr.Login, string(hash))
	if err != nil {
		return 0, err
	}
	return id, nil
}
