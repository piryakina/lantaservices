package usecase

type SP struct {
	ID          int64  `json:"id"`
	NameCompany string `json:"name_company,omitempty"`
	Login       string `json:"login,omitempty"`
	Password    string `json:"password,omitempty"`
	Email       string `json:"email,omitempty"`
	Phone       string `json:"phone,omitempty"`
}
