package models

import "time"

type DateEvent struct {
	ID          int       `json:"id" db:"id"`
	Person1Name string    `json:"person1_name" db:"person1_name"`
	Person2Name string    `json:"person2_name" db:"person2_name"`
	Date        time.Time `json:"date" db:"date"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
