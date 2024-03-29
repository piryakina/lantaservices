package usecase

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"lantaservice/entities"
	"lantaservice/storage"
)

//SignUpUser - регистриация пользователя
func SignUpUser(ctx context.Context, usr *entities.User) (int64, error) { //registration  (s *ServiceUser)
	hash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	//id, err := s.UserRepository.SignUpStorage(usr.Login, string(hash))
	usr.Password = string(hash)
	id, err := storage.AddUser(ctx, usr)
	if err != nil {
		return 0, err
	}
	return id, nil
}
//GetRoleUserById - получени роли пользователя
func GetRoleUserById(ctx context.Context, id int64) (string, string, error) {
	role, name, err := storage.GetUserRoleById(ctx, id)
	if err != nil {
		return "", "", err
	}
	return role, name, nil
}
//GetUserRoles - получение списка ролей
func GetUserRoles(ctx context.Context) ([]string, error) {
	role, err := storage.GetRoles(ctx)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func CheckUserLogin(ctx context.Context, login string) (bool, error) {
	return storage.CheckUserLoginStorage(ctx, login)
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
