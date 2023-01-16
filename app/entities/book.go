package entities

import "time"

type Book struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Name      string
	Author    string
	CreatedBy string
}
