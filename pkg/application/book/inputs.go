package book

import domain "github.com/delgerskhn/gosetup/pkg/domain/entities"

type CreateBookInput struct {
	Name   string `validate:"required,min=4,max=15"`
	Author string `binding:"required"`
}

func (b CreateBookInput) ToEntity() domain.Book {
	return domain.Book{
		Name:   b.Name,
		Author: b.Author,
	}
}
