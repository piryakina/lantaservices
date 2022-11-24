package storage

import (
	"context"
	"database/sql"
	"lantaservice/usecase"
)

type SPDB struct {
	ID          int64          `db:"id"`
	NameCompany sql.NullString `db:"name_company"`
	Email       sql.NullString `db:"email"`
	Phone       sql.NullString `db:"phone"`
	Login       sql.NullString `db:"login"`
	Password    sql.NullString `db:"password""`
}

func FromSPDB(p *SPDB) *usecase.SP {
	var n string
	if p.NameCompany.Valid {
		n = p.NameCompany.String
	}
	var login string
	if p.Login.Valid {
		login = p.Login.String
	}
	var pwd string
	if p.Password.Valid {
		pwd = p.Password.String
	}
	var mail string
	if p.Email.Valid {
		mail = p.Email.String
	}
	var phone string
	if p.Phone.Valid {
		phone = p.Phone.String
	}

	return &usecase.SP{
		ID:          p.ID,
		Login:       login,
		Password:    pwd,
		Email:       mail,
		Phone:       phone,
		NameCompany: n,
	}
}

// AddSP - add sp to db
func AddSP(ctx context.Context, sp *usecase.SP) (int64, error) {
	db, err := GetDB()
	if err != nil {
		return 0, err
	}
	query := "INSERT INTO sp (name_company, email, phone, login, password) VALUES ($1,$2,$3,$4,$5) returning id"
	var id int64
	row := db.QueryRowContext(ctx, query, sp.NameCompany, sp.Email, sp.Phone, sp.Login, sp.Password)
	if err = row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// GetSPById - get sp by id
func GetSPById(ctx context.Context, id int64) (*usecase.SP, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	query := "SELECT * from sp WHERE id=$1"
	row := db.QueryRowContext(ctx, query, id)
	var sp SPDB
	if err = row.Scan(&sp.ID, &sp.NameCompany, &sp.Email, &sp.Phone, &sp.Login, &sp.Password); err != nil {
		return nil, err
	}
	var partner *usecase.SP
	partner = FromSPDB(&sp)
	return partner, nil
}
