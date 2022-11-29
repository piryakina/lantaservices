package storage

import (
	"context"
	"database/sql"
	"lantaservice/entities"
)

type UserDB struct {
	ID       int64          `db:"id"`
	FIO      sql.NullString `db:"fio"`
	Login    sql.NullString `db:"login"`
	Password sql.NullString `db:"password"`
	Email    sql.NullString `db:"email"`
	Phone    sql.NullString `db:"phone,omitempty"`
	Role     string         `db:"role"`
}

func FromUserDB(p *UserDB) *entities.User {
	var fio string
	if p.FIO.Valid {
		fio = p.FIO.String
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

	return &entities.User{
		ID:       p.ID,
		FIO:      fio,
		Login:    login,
		Password: pwd,
		Email:    mail,
		Phone:    phone,
		Role:     p.Role,
	}
}

//const connStr = "user=postgres password=Wt2H1aqF dbname=lanta sslmode=disable"
//
////func GetDB() (*sql.DB, error) {
////	db, err := sql.Open("postgres", connStr)
////	if err != nil {
////		return nil, err
////	}
////	return db, nil
////}

// AddUser  - add user to db
func AddUser(ctx context.Context, usr *entities.User) (int64, error) {
	//(s * Storage)
	db, err := GetDB()
	if err != nil {
		return 0, err
	}
	var idRole int64
	query := "SELECT id from \"role\" WHERE role=$1"
	if usr.Role != "" {
		row := db.QueryRowContext(ctx, query, usr.Role)
		if err = row.Scan(&idRole); err != nil {
			return 0, err
		}
	} else {
		idRole = 1
	}
	query = "INSERT INTO \"user\" (fio, login, email, phone, password, \"role\") VALUES ($1,$2,$3,$4,$5,$6) returning id"
	var id int64
	row := db.QueryRowContext(ctx, query, usr.FIO, usr.Login, usr.Email, usr.Phone, usr.Password, idRole)
	if err = row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// GetUserById - get user by id
func GetUserById(ctx context.Context, id int64) (*entities.User, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	query := "SELECT * from \"user\" WHERE id=$1"
	row := db.QueryRowContext(ctx, query, id)
	var usr UserDB
	var idRole int64
	if err = row.Scan(&usr.ID, &usr.FIO, &usr.Email, &usr.Phone, &usr.Password, &idRole); err != nil {
		return nil, err
	}
	query = "SELECT role from role where id=$1"
	row = db.QueryRowContext(ctx, query, id)
	if err = row.Scan(&usr.Role); err != nil {
		return nil, err
	}
	var user *entities.User
	user = FromUserDB(&usr)
	return user, nil
}

func LoginUserStorage(ctx context.Context, usr string) (int64, string, error) {
	db, err := GetDB()
	if err != nil {
		return 0, "", err
	}
	query := "SELECT id,password from \"user\" WHERE login=$1"
	var pwd string
	var id int64
	row := db.QueryRowContext(ctx, query, usr)

	if err = row.Scan(&id, &pwd); err != nil {
		return 0, "", err
	}
	return id, pwd, nil
}

//func SignUpStorage(usr string, pwd string) (int64, error) { //registration
//	db, err := GetDB()
//	if err != nil {
//		return 0, err
//	}
//	query := "INSERT INTO \"user\" (login, password) VALUES ($1, $2) returning id"
//	var id int64
//	row := db.QueryRow(query, usr, pwd)
//	if row.Err() != nil {
//		return 0, row.Err()
//	}
//	err = row.Scan(&id)
//	if err != nil {
//		return 0, err
//	}
//	return id, nil
//}
