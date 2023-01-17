package domain

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

func (b Book) ChangeAuthor(author string) Book {
	return Book{
		ID:        b.ID,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
		DeletedAt: b.DeletedAt,
		Name:      b.Name,
		Author:    author,
		CreatedBy: b.CreatedBy,
	}
}
