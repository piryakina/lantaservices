package usecase

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"lantaservice/entities"
)

type UserServer interface {
	SignUpUser(ctx context.Context, usr *entities.User) (int64, error)
	SignUpSP(ctx context.Context, usr *entities.SP) (int64, error)
}
type ServiceUser struct {
	UserRepository entities.UserRepository
}

func (s *ServiceUser) SignUpUser(ctx context.Context, usr *entities.User) (int64, error) { //registration
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

type Hash struct{}

//// Generate a salted hash for the input string
//func Generate(s string) (string, error) {
//	saltedBytes := []byte(s)
//	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
//	if err != nil {
//		return "", err
//	}
//
//	hash := string(hashedBytes[:])
//	return hash, nil
//}

// Compare string to generated hash
func Compare(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}
