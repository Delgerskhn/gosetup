package repository

import (
	"context"

	"github.com/delgerskhn/gosetup/app/entities"
	"github.com/delgerskhn/gosetup/app/gorm"
	"github.com/delgerskhn/gosetup/infrastructure/persistence"
	repo "github.com/delgerskhn/gosetup/infrastructure/persistence/repository"
)

var repository = repo.NewRepository[gorm.BookGorm, entities.Book](db)

var db = persistence.GetDB()

func GetBooks() ([]entities.Book, error) {
	return repository.FindAll(context.Background())

}

func CreateBook(book entities.Book) (entities.Book, error) {
	err := repository.Insert(context.Background(), &book)
	return book, err
}

func GetBook(id string) (entities.Book, error) {
	return repository.FindByID(context.Background(), id)
}

func DeleteBook(id string) error {
	return repository.DeleteById(context.Background(), id)
}

func UpdateBook(book entities.Book) (entities.Book, error) {
	return book, repository.Update(context.Background(), &book)
}
