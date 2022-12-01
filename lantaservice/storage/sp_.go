package storage

import (
	"context"
	"database/sql"
	"lantaservice/entities"
	"time"
)

type SPDB struct {
	ID          int64          `db:"id"`
	NameCompany sql.NullString `db:"name_company"`
	Email       sql.NullString `db:"email"`
	Phone       sql.NullString `db:"phone"`
	Login       sql.NullString `db:"login"`
	Password    sql.NullString `db:"password"`
}

type SpPeriodDB struct {
	ID       int64          `db:"id"`
	SP       sql.NullString `db:"sp"`
	Period   sql.NullString `db:"period"`
	Quality  sql.NullString `db:"quality"`
	Invoice  sql.NullString `db:"invoice"`
	Vehicles int64          `db:"vehicle_service"`
}

func FromSPDB(p *SPDB) *entities.SP {
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

	return &entities.SP{
		ID:          p.ID,
		Login:       login,
		Password:    pwd,
		Email:       mail,
		Phone:       phone,
		NameCompany: n,
	}
}

func FromSPPeriodDB(p *SpPeriodDB) *entities.SpPeriod {
	var sp string
	if p.SP.Valid {
		sp = p.SP.String
	}
	var per string
	if p.Period.Valid {
		per = p.Period.String
	}
	var qu string
	if p.Quality.Valid {
		qu = p.Quality.String
	}
	var inv string
	if p.Invoice.Valid {
		inv = p.Invoice.String
	}
	return &entities.SpPeriod{
		ID:      p.ID,
		Sp:      sp,
		Period:  per,
		Quality: qu,
		Invoice: entities.InvoiceFile{
			Filename: inv,
			Date:     time.Time{},
		},
		Vehicle: p.Vehicles,
	}
}

// AddSP - add sp to db
func AddSP(ctx context.Context, sp *entities.SP) (int64, error) {
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
func GetSPById(ctx context.Context, id int64) (*entities.SP, error) {
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
	var partner *entities.SP
	partner = FromSPDB(&sp)
	return partner, nil
}

//func LoginSpStorage(ctx context.Context, usr string) (int64, string, error) {
//	db, err := GetDB()
//	if err != nil {
//		return 0, "", err
//	}
//	query := "SELECT id,password from sp WHERE login=$1"
//	var pwd string
//	var id int64
//	row := db.QueryRowContext(ctx, query, usr)
//
//	if err = row.Scan(&id, &pwd); err != nil {
//		return 0, "", err
//	}
//	return id, pwd, nil
//}

func GetDataSpPeriodStorage(ctx context.Context, login string, date time.Time) (*entities.SpPeriod, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	var idSp, idPeriod int64
	var temp SpPeriodDB
	query := "SELECT id,name from \"user\" WHERE login=$1"
	row := db.QueryRowContext(ctx, query, login)
	if err = row.Scan(&idSp, &temp.SP); err != nil {
		return nil, err
	}
	query = "SELECT id,title FROM period WHERE $1 between date_from and date_to"
	row = db.QueryRowContext(ctx, query, date)
	if err = row.Scan(&idPeriod, &temp.Period); err != nil {
		return nil, err
	}
	query = "SELECT t1.vehicle_service,t1.quality, (select filename as invoice from invoice_file as t2 where t2.id=t1.invoice) from sp_period as t1 where t1.sp=$1 and t1.period=$2"
	row = db.QueryRowContext(ctx, query, idSp, idPeriod)
	if err != nil {
		return nil, err
	}
	var res *entities.SpPeriod

	if err = row.Scan(&temp.Vehicles, &temp.Quality, &temp.Invoice); err != nil {
		return nil, err
	} //res = append(res)
	res = FromSPPeriodDB(&temp)
	return res, nil
}

func AddDataSpPeriodStorage(ctx context.Context, data *entities.SpPeriod) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	query := "SELECT id from \"user\" WHERE name=$1"
	row := db.QueryRowContext(ctx, query, data.Sp)
	var idSP, idPeriod, idInvoice int64
	if err = row.Scan(&idSP); err != nil {
		return err
	}
	query = "SELECT id FROM period WHERE title=$1 "
	row = db.QueryRowContext(ctx, query, data.Period)
	if err = row.Scan(&idPeriod); err != nil {
		return err
	}
	//query = "SELECT t1.vehicle_service,t1.quality, (select filename as invoice from invoice_file as t2 where t2.id=t1.invoice) from sp_period as t1 where t1.sp=$1 and t1.period=$2"
	if data.Invoice.Filename != "" {
		query = "insert into invoice_file (filename, path, date) VALUES ($1,$2,$3) returning id"
		data.Invoice.Date = time.Now()
		row = db.QueryRowContext(ctx, query, data.Invoice.Filename, data.Invoice.Path, data.Invoice.Date)
		if err = row.Scan(&idInvoice); err != nil {
			return err
		}
	}
	query = "insert into sp_period (sp,period,invoice, vehicle_service,quality) VALUES ($1,$2,$3,$4,$5)"
	row = db.QueryRowContext(ctx, query, idSP, idPeriod, idInvoice, data.Vehicle, data.Quality)
	if err != nil {
		return err
	}
	return nil
}

//func GetSpNameByID(ctx context.Context, id int64) (string, error) {
//	db, err := GetDB()
//	if err != nil {
//		return "", err
//	}
//	query := "SELECT name_company from sp WHERE id=$1"
//	row := db.QueryRowContext(ctx, query, id)
//	var name string
//	if err = row.Scan(&name); err != nil {
//		return "", err
//	}
//	return name, err
//}
