package storage

import (
	"context"
	"database/sql"
	"lantaservice/entities"
)

type NewsDB struct {
	Id     int64              `db:"id"`
	Title  sql.NullString     `db:"title"`
	Text   sql.NullString     `db:"text"`
	Date   sql.NullString     `db:"date"`
	Attach []*entities.Attach `db:"attach"`
}

type AttachDB struct {
	Id       int64  `db:"id"`
	Path     string `db:"path"`
	Filename string `db:"filename"`
	NewsId   int64  `db:"newsId"`
}

func fromNewsDB(p NewsDB) *entities.News {
	var title string
	if p.Title.Valid {
		title = p.Title.String
	}
	var text string
	if p.Text.Valid {
		text = p.Text.String
	}
	//layout := "2006-01-02 15:04:05Z" //todo yyyy-mm-dd
	//var date time.Time
	//var d time.Time
	//if p.Date.Valid {
	//	d, err := time.Parse(time.RFC3339Nano, p.Date.String)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(d)
	//}
	var date string
	if p.Date.Valid {
		date = p.Date.String
	}
	return &entities.News{
		Id:          p.Id,
		Title:       title,
		Text:        text,
		Date:        date,
		Attachments: p.Attach,
	}
}
func FromAttachDB(p AttachDB) *entities.Attach {
	return &entities.Attach{
		Id:       p.Id,
		Path:     p.Path,
		Filename: p.Filename,
		NewsId:   p.NewsId,
	}
}

func GetNewsStorage(ctx context.Context) ([]*entities.News, error) {
	db := GetDB()
	query := "SELECT title,text,\"date\" from news"
	rows, err := db.QueryContext(ctx, query)
	var news []*entities.News
	for rows.Next() {
		var c NewsDB
		if err = rows.Scan(&c.Id, &c.Title, &c.Text, &c.Date); err != nil {
			return nil, err
		}
		query = "select id,path,filename from attachment where news_id=$1"
		res, err := db.QueryContext(ctx, query, c.Id)
		if err != nil {
			return nil, err
		}
		for res.Next() {
			var img AttachDB
			if err = res.Scan(&img.Id, &img.Path, &img.Filename); err != nil {
				return nil, err
			}
			c.Attach = append(c.Attach, FromAttachDB(img))
		}
		news = append(news, fromNewsDB(c))
	}
	return news, nil
}

func AddNewsStorage(ctx context.Context, p *entities.News) (int64, error) {
	db := GetDB()
	query := "INSERT INTO news (title,text,\"date\") VALUES  ($1,$2,$3) returning id"
	title := ToNullString(p.Title)
	text := ToNullString(p.Text)
	date := ToNullString(p.Date)
	row := db.QueryRowContext(ctx, query, title, text, date)
	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
