package entities

import "time"

type SpPeriod struct {
	ID      int64         `json:"id,omitempty"`
	Sp      string        `json:"sp,omitempty"`
	Period  string        `json:"period,omitempty"`
	Quality string        `json:"quality,omitempty"`
	Invoice []InvoiceFile `json:"invoice,omitempty"`
	Vehicle int64         `json:"vehicle,omitempty"`
	Billing []BillingFile `json:"billing,omitempty"`
}

type InvoiceFile struct {
	ID       int64     `json:"id,omitempty"`
	Filename string    `json:"filename,omitempty"`
	Path     string    `json:"path,omitempty"`
	Date     time.Time `json:"date,omitempty"`
}

type BillingFile struct {
	ID       int64     `json:"id,omitempty"`
	Filename string    `json:"filename,omitempty"`
	Path     string    `json:"path,omitempty"`
	Date     time.Time `json:"date,omitempty"`
	Status   string    `json:"status,omitempty"`
}
type CommentFile struct {
	ID      int64  `json:"id,omitempty"`
	Comment string `json:"comments,omitempty"`
}
