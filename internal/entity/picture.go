package entity

import "time"

type Picture struct {
	ID          int64  `json:"id" db:"id"`
	ImageURL    string `json:"image_url" db:"image_url"`
	PictureDate time.Time
}
