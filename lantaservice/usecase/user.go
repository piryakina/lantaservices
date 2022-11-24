package usecase

type User struct {
	ID       int64  `json:"id"`
	FIO      string `json:"fio"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone,omitempty"`
	Role     string `json:"role"`
}
