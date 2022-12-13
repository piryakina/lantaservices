package entities

type News struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Text        string    `json:"text"`
	Date        string    `json:"date"`
	Attachments []*Attach `json:"attachments,omitempty"`
}

type Attach struct {
	Id       int64  `json:"id"`
	Path     string `json:"path"`
	Filename string `json:"filename"`
	NewsId   int64  `json:"news_id,omitempty"`
}

//type NewsRepository interface {
//	GetNewsStorage(ctx context.Context) ([]*News, error)
//	AddNewsStorage(ctx context.Context, p *News) error
//}
