package inputs

import "github.com/delgerskhn/gosetup/app/entities"

type CreateBookInput struct {
	Name   string `json:"name" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func (b CreateBookInput) ToEntity() entities.Book {
	return entities.Book{
		Name:   b.Name,
		Author: b.Author,
	}
}
