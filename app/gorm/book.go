package gorm

import (
	"github.com/delgerskhn/gosetup/app/entities"
	"gorm.io/gorm"
)

type BookGorm struct {
	gorm.Model
	Name      string
	Author    string
	CreatedBy string
}

func (g BookGorm) ToEntity() entities.Book {
	return entities.Book{
		ID:        g.ID,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
		DeletedAt: g.DeletedAt.Time,
		Name:      g.Name,
	}
}

// FromEntity respects the gorm_generics.GormModel interface
// Creates new GORM model from Entity.
func (g BookGorm) FromEntity(book entities.Book) interface{} {
	return BookGorm{
		Model: gorm.Model{
			ID:        book.ID,
			CreatedAt: book.CreatedAt,
			UpdatedAt: book.UpdatedAt,
			DeletedAt: gorm.DeletedAt{Time: book.DeletedAt},
		},
		Name:      book.Name,
		Author:    book.Author,
		CreatedBy: book.CreatedBy,
	}
}
