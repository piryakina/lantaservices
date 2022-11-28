package entities

import (
	"context"
)

type UserLogin struct {
	Id       int64
	Login    string
	Password string
}

type User struct {
	ID       int64  `json:"id"`
	FIO      string `json:"fio"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone,omitempty"`
	Role     string `json:"role"`
}

type SP struct {
	ID          int64  `json:"id"`
	NameCompany string `json:"name_company,omitempty"`
	Login       string `json:"login,omitempty"`
	Password    string `json:"password,omitempty"`
	Email       string `json:"email,omitempty"`
	Phone       string `json:"phone,omitempty"`
}

type UserRepository interface {
	SignUpStorage(ctx context.Context, usr string, pwd string) (int64, error)
	LoginUserStorage(ctx context.Context, usr string) (int64, string, error)
	GetUserById(ctx context.Context, id int64) (*User, error)
	AddUser(ctx context.Context, usr *User) (int64, error)

	AddSP(ctx context.Context, sp *SP) (int64, error)
	LoginSpStorage(ctx context.Context, usr string) (int64, string, error)
	GetSPById(ctx context.Context, id int64) (*SP, error)
}
