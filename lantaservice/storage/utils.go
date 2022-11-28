package storage

import (
	"database/sql"
	"time"
)

func ToNullInt(val int64) sql.NullInt64 {
	if val == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{Int64: val, Valid: true}
}

func ToNullString(val string) sql.NullString {
	if val == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: val, Valid: true}
}

func ToNullTime(val time.Time) sql.NullTime {
	if val.IsZero() {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  val,
		Valid: true,
	}
}
