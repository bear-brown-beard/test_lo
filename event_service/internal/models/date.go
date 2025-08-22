package models

import "time"

type DateEvent struct {
	ID          int       `json:"id" db:"id"`
	Date        time.Time `json:"date" db:"date"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
