package entities

type News struct {
	Id          int64     `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Text        string    `json:"text,omitempty"`
	Date        string    `json:"date,omitempty"`
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
