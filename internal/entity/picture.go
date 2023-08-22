package entity

import "time"

type Picture struct {
	ID   int64  `json:"id" db:"id"`
	URL  string `json:"url" db:"url"`
	Date time.Time
}
