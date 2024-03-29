package storage

import (
	"context"
	"database/sql"
	"fmt"
	"lantaservice/entities"
	"log"
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
	Vehicles sql.NullInt64  `db:"vehicle_service"`
}
type BillingFileDB struct {
	ID       int64          `db:"id"`
	Filename sql.NullString `db:"filename"`
	Path     sql.NullString `db:"path"`
	Date     string         `db:"date"`
	Status   sql.NullString `db:"status"`
	Comments sql.NullString `db:"comments"`
}

type InvoiceFileDB struct {
	ID       int64          `db:"id"`
	Filename sql.NullString `db:"filename"`
	Path     sql.NullString `db:"path"`
	Date     string         `db:"date"`
}
type SLAFileDB struct {
	ID       int64          `db:"id"`
	Filename sql.NullString `db:"filename"`
	Path     sql.NullString `db:"path"`
	//Date     string         `db:"date"`
	USP      sql.NullString `db:"usp"`
	SpPeriod int64          `db:"sp_period"`
	IsAgreed sql.NullBool   `db:"is_agreed"`
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
func fromFileDB(p BillingFileDB) *entities.BillingFile {
	var filename string
	if p.Filename.Valid {
		filename = p.Filename.String
	}
	var comment string
	if p.Comments.Valid {
		comment = p.Comments.String
	}
	var st string
	if p.Status.Valid {
		st = p.Status.String
	}
	var path string
	if p.Path.Valid {
		path = p.Path.String
	}
	//layout := "2006-01-02" //todo yyyy-mm-dd
	var date time.Time
	fmt.Println(p.Date)
	date, err := time.Parse("2006-01-02T15:04:05Z", p.Date)
	fmt.Println(date)
	//date = date.Format("2006-01-02")
	if err != nil {
		log.Fatal(err)
	}
	return &entities.BillingFile{
		ID:       p.ID,
		Filename: filename,
		Path:     path,
		Date:     date,
		Status:   st,
		Comments: comment,
	}
}

func fromInvoiceDB(p InvoiceFileDB) *entities.InvoiceFile {
	var filename string
	if p.Filename.Valid {
		filename = p.Filename.String
	}
	var path string
	if p.Path.Valid {
		path = p.Path.String
	}
	//layout := "2006-01-02" //todo yyyy-mm-dd
	var date time.Time
	fmt.Println(p.Date)
	date, err := time.Parse("2006-01-02T15:04:05Z", p.Date)
	fmt.Println(date)
	//date = date.Format("2006-01-02")
	if err != nil {
		log.Fatal(err)
	}
	return &entities.InvoiceFile{
		ID:       p.ID,
		Filename: filename,
		Path:     path,
		Date:     date,
	}
}

func fromSLADB(p SLAFileDB) *entities.SLAFile {
	var filename string
	if p.Filename.Valid {
		filename = p.Filename.String
	}
	var path string
	if p.Path.Valid {
		path = p.Path.String
	}
	var usp string
	if p.USP.Valid {
		usp = p.USP.String
	}
	//layout := "2006-01-02" //todo yyyy-mm-dd
	//var date time.Time
	//fmt.Println(p.Date)
	//date, err := time.Parse("2006-01-02T15:04:05Z", p.Date)
	//fmt.Println(date)
	////date = date.Format("2006-01-02")
	//if err != nil {
	//	log.Fatal(err)
	//}
	var agree bool
	if p.IsAgreed.Valid {
		agree = p.IsAgreed.Bool
	}
	return &entities.SLAFile{
		ID:       p.ID,
		Filename: filename,
		Path:     path,
		USP:      usp,
		SpPeriod: p.SpPeriod,
		IsAgreed: &agree,
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
	var vech int64
	if p.Vehicles.Valid {
		vech = p.Vehicles.Int64
	}
	return &entities.SpPeriod{
		ID:      p.ID,
		Sp:      sp,
		Period:  per,
		Quality: qu,
		//Invoice: entities.InvoiceFile{
		//	Filename: inv,
		//	Date:     time.Time{},
		//},
		Vehicle: vech,
	}
}

// AddSP - add sp to db
func AddSP(ctx context.Context, sp *entities.SP) (int64, error) {
	db := GetDB()
	query := "INSERT INTO sp (name_company, email, phone, login, password) VALUES ($1,$2,$3,$4,$5) returning id"
	var id int64
	row := db.QueryRowContext(ctx, query, sp.NameCompany, sp.Email, sp.Phone, sp.Login, sp.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// GetSPById - get sp by id
func GetSPById(ctx context.Context, id int64) (*entities.SP, error) {
	db := GetDB()
	query := "SELECT * from sp WHERE id=$1"
	row := db.QueryRowContext(ctx, query, id)
	var sp SPDB
	if err := row.Scan(&sp.ID, &sp.NameCompany, &sp.Email, &sp.Phone, &sp.Login, &sp.Password); err != nil {
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
	db := GetDB()
	var idSp, idPeriod int64
	var temp SpPeriodDB
	query := "SELECT id,name from \"user\" WHERE login=$1"
	row := db.QueryRowContext(ctx, query, login)
	if err := row.Scan(&idSp, &temp.SP); err != nil {
		return nil, err
	}
	query = "SELECT id,title FROM period WHERE $1 between date_from and date_to"
	row = db.QueryRowContext(ctx, query, date)
	err := row.Scan(&idPeriod, &temp.Period)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		log.Fatalf("query error: %v\n", err)
		return nil, err
	}
	query = "SELECT t1.id,t1.vehicle_service,t1.quality from sp_period as t1 where t1.sp=$1 and t1.period=$2"
	row = db.QueryRowContext(ctx, query, idSp, idPeriod)
	var res *entities.SpPeriod
	if err := row.Scan(&temp.ID, &temp.Vehicles, &temp.Quality); err != nil {
		return nil, err
	} //res = append(res)
	query = "select b.id, b.filename,b.path,b.date, (select d.status_name from docs_status as d where d.id=b.status) as status, comments from billing_file as b where b.sp_period_id=$1"
	rows, err := db.QueryContext(ctx, query, temp.ID)
	if err != nil {
		return nil, err
	}
	var billings []entities.BillingFile
	for rows.Next() {
		var file BillingFileDB
		if err = rows.Scan(&file.ID, &file.Filename, &file.Path, &file.Date, &file.Status, &file.Comments); err != nil {
			return nil, err
		}
		//fmt.Println(file.ID)
		billings = append(billings, *fromFileDB(file))
	}
	defer rows.Close()
	query = "select v.id,v.filename, v.path, v.date from invoice_file as v where sp_period_id=$1"
	rows2, err := db.QueryContext(ctx, query, temp.ID)
	var invoices []entities.InvoiceFile
	for rows2.Next() {
		var file entities.InvoiceFile
		if err = rows2.Scan(&file.ID, &file.Filename, &file.Path, &file.Date); err != nil {
			return nil, err
		}
		invoices = append(invoices, file)
	}
	query = "SELECT id, filename, path, is_agreed from sla_file where sp_period=$1"
	var rows3 *sql.Rows
	rows3, err = db.QueryContext(ctx, query, temp.ID)
	if err != nil {
		return nil, err
	}
	var t SLAFileDB
	for rows3.Next() {
		if err = rows3.Scan(&t.ID, &t.Filename, &t.Path, &t.IsAgreed); err != nil {
			return nil, err
		}
		//temp.SLA = *fromSLADB(t)
	}
	defer rows2.Close()
	res = FromSPPeriodDB(&temp)
	res.Billing = billings
	res.Invoice = invoices
	res.SLA = *fromSLADB(t)
	return res, nil
}

func AddDataSpPeriodStorage(ctx context.Context, data *entities.SpPeriod) error {
	db := GetDB()

	query := "SELECT id from \"user\" WHERE login=$1"
	row := db.QueryRowContext(ctx, query, data.Sp)
	var idSP, idPeriod, idSpPeriod int64
	if err := row.Scan(&idSP); err != nil {
		return err
	}
	query = "SELECT id FROM period WHERE title=$1 "
	row = db.QueryRowContext(ctx, query, data.Period)
	if err := row.Scan(&idPeriod); err != nil {
		return err
	}
	query = "insert into sp_period (sp,period,vehicle_service,quality) VALUES ($1,$2,$3,$4) returning id"
	row = db.QueryRowContext(ctx, query, idSP, idPeriod, data.Vehicle, data.Quality)
	if err := row.Scan(&idSpPeriod); err != nil {
		return err
	}
	if len(data.Billing) != 0 {
		for i := 0; i < len(data.Billing); i++ {
			if data.Billing[i].Filename != "" {
				query = "insert into billing_file (filename,path,date,status, sp_period_id) values ($1,$2,$3,$4,$5) returning id"
				var bId int64
				err := db.QueryRowContext(ctx, query, data.Billing[i].Filename, data.Billing[i].Path, time.Now(), 1, idSpPeriod).Scan(&bId)
				if err != nil {
					log.Fatal(err)
					return err
				}
			}
		}
	}
	if len(data.Invoice) != 0 {
		for i := 0; i < len(data.Invoice); i++ {
			if data.Billing[i].Filename != "" {
				query = "insert into invoice_file (filename,path,date, sp_period_id) values ($1,$2,$3,$4) returning id"
				var id int64
				err := db.QueryRowContext(ctx, query, data.Billing[i].Filename, data.Billing[i].Path, time.Now(), idSpPeriod).Scan(&id)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

//	func GetSpNameByID(ctx context.Context, id int64) (string, error) {
//		db, err := GetDB()
//		if err != nil {
//			return "", err
//		}
//		query := "SELECT name_company from sp WHERE id=$1"
//		row := db.QueryRowContext(ctx, query, id)
//		var name string
//		if err = row.Scan(&name); err != nil {
//			return "", err
//		}
//		return name, err
//	}
func GetDataPeriodStorage(ctx context.Context, idPeriod int64) ([]*entities.SpPeriod, error) { //todo get rows sp_period
	//--SELECT sp.id,(select t2.name from "user" as t2 where id=sp.sp) as name, (select t3.title from period as t3 where id=sp.period) as period from sp_period as sp
	//	SELECT filename, path, status from billing_file where sp_period_id=2
	db := GetDB()
	var res []*entities.SpPeriod
	query := "SELECT sp.id,(select t2.name from \"user\" as t2 where id=sp.sp) as sp, (select t3.title from period as t3 where id=sp.period) as period from sp_period as sp where period=$1"
	rows, err := db.QueryContext(ctx, query, idPeriod)
	if err != nil {
		return nil, err
	}
	var rows2, rows3 *sql.Rows
	for rows.Next() {
		var temp entities.SpPeriod
		if err = rows.Scan(&temp.ID, &temp.Sp, &temp.Period); err != nil {
			return nil, err
		}
		query = "SELECT id,filename, path, status,date, comments from billing_file where sp_period_id=$1"
		rows2, err = db.QueryContext(ctx, query, temp.ID)
		if err != nil {
			return nil, err
		}
		for rows2.Next() {
			var t BillingFileDB
			if err = rows2.Scan(&t.ID, &t.Filename, &t.Path, &t.Status, &t.Date, &t.Comments); err != nil {
				return nil, err
			}
			temp.Billing = append(temp.Billing, *fromFileDB(t))
		}
		query = "SELECT id,filename, path,date from invoice_file where sp_period_id=$1"
		rows2, err = db.QueryContext(ctx, query, temp.ID)
		if err != nil {
			return nil, err
		}
		for rows2.Next() {
			var t InvoiceFileDB
			if err = rows2.Scan(&t.ID, &t.Filename, &t.Path, &t.Date); err != nil {
				return nil, err
			}
			temp.Invoice = append(temp.Invoice, *fromInvoiceDB(t))
		}
		query = "SELECT id, filename, path, is_agreed from sla_file where sp_period=$1"
		rows3, err = db.QueryContext(ctx, query, temp.ID)
		if err != nil {
			return nil, err
		}
		for rows3.Next() {
			var t SLAFileDB
			if err = rows3.Scan(&t.ID, &t.Filename, &t.Path, &t.IsAgreed); err != nil {
				return nil, err
			}
			temp.SLA = *fromSLADB(t)
			//fmt.Println(temp.SLA.IsAgreed)
		}
		res = append(res, &temp)

	}
	defer rows3.Close()
	defer rows2.Close()
	defer rows.Close()
	return res, nil
}

func SetApproveSLA(ctx context.Context, approve bool, id int64) error {
	query := "Update sla_file set is_agreed=$1 where id=$2"
	db := GetDB()
	_, err := db.ExecContext(ctx, query, approve, id)
	if err != nil {
		return err
	}
	return nil
}
