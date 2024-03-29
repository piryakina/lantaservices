package storage

import (
	"context"
	"database/sql"
	"lantaservice/entities"
)

type UserDB struct {
	ID       int64          `db:"id"`
	Name     sql.NullString `db:"name"`
	Login    sql.NullString `db:"login"`
	Password sql.NullString `db:"password"`
	Email    sql.NullString `db:"email"`
	Phone    sql.NullString `db:"phone,omitempty"`
	Role     string         `db:"role"`
}

func FromUserDB(p *UserDB) *entities.User {
	var name string
	if p.Name.Valid {
		name = p.Name.String
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
		Name:     name,
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
	db := GetDB()
	var idRole int64
	query := "SELECT id from \"role\" WHERE role=$1"
	if usr.Role != "" {
		row := db.QueryRowContext(ctx, query, usr.Role)
		if err := row.Scan(&idRole); err != nil {
			return 0, err
		}
	} else {
		idRole = 1
	}
	query = "INSERT INTO \"user\" (name, login, email, phone, password, \"role\") VALUES ($1,$2,$3,$4,$5,$6) returning id"
	var id int64
	row := db.QueryRowContext(ctx, query, usr.Name, usr.Login, usr.Email, usr.Phone, usr.Password, idRole)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// GetUserById - get user by id
func GetUserById(ctx context.Context, id int64) (*entities.User, error) {
	db := GetDB()
	query := "SELECT * from \"user\" WHERE id=$1"
	row := db.QueryRowContext(ctx, query, id)
	var usr UserDB
	var idRole int64
	if err := row.Scan(&usr.ID, &usr.Name, &usr.Email, &usr.Phone, &usr.Password, &idRole); err != nil {
		return nil, err
	}
	query = "SELECT role from role where id=$1"
	row = db.QueryRowContext(ctx, query, id)
	if err := row.Scan(&usr.Role); err != nil {
		return nil, err
	}
	var user *entities.User
	user = FromUserDB(&usr)
	return user, nil
}

func LoginUserStorage(ctx context.Context, usr string) (int64, string, string, string, error) {
	db := GetDB()
	query := "SELECT id,password,role,name from \"user\" WHERE login=$1"
	var pwd, role, name string
	var id, roleId int64
	row := db.QueryRowContext(ctx, query, usr)

	if err := row.Scan(&id, &pwd, &roleId, &name); err != nil {
		return 0, "", "", "", err
	}
	query = "SELECT role from role where id=$1"
	row = db.QueryRowContext(ctx, query, roleId)
	if err := row.Scan(&role); err != nil {
		return 0, "", "", "", err
	}
	return id, pwd, role, name, nil
}

// func SignUpStorage(usr string, pwd string) (int64, error) { //registration
//
//		db, err := GetDB()
//		if err != nil {
//			return 0, err
//		}
//		query := "INSERT INTO \"user\" (login, password) VALUES ($1, $2) returning id"
//		var id int64
//		row := db.QueryRow(query, usr, pwd)
//		if row.Err() != nil {
//			return 0, row.Err()
//		}
//		err = row.Scan(&id)
//		if err != nil {
//			return 0, err
//		}
//		return id, nil
//	}
//
// GetUserById - get user by id
func GetUserRoleById(ctx context.Context, id int64) (string, string, error) {
	db := GetDB()
	if id != 0 {
		query := "SELECT role,name from \"user\" WHERE id=$1"
		row := db.QueryRowContext(ctx, query, id)
		var idRole int64
		var role, name string
		if err := row.Scan(&idRole, &name); err != nil {
			return "", "", err
		}
		query = "SELECT role from role where id=$1"
		row = db.QueryRowContext(ctx, query, idRole)
		if err := row.Scan(&role); err != nil {
			return "", "", err
		}
		return role, name, nil

	}
	return "", "", nil
}
func GetRoles(ctx context.Context) ([]string, error) {
	db := GetDB()
	query := "SELECT role from role"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var roles []string
	for rows.Next() {
		var r string
		if err = rows.Scan(&r); err != nil {
			return nil, err
		}
		roles = append(roles, r)
	}
	return roles, nil
}
func CheckUserLoginStorage(ctx context.Context, login string) (bool, error) {
	db := GetDB()
	query := "select id from \"user\" where login=$1"
	var id int64
	err := db.QueryRowContext(ctx, query, login).Scan(&id)
	if err != nil {
		return false, err
	}
	if id != 0 {
		return false, nil
	} else {
		return true, nil
	}
}
