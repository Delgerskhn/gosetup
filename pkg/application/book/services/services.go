package services

import (
	"context"

	domain "github.com/delgerskhn/gosetup/pkg/domain/entities"
	"github.com/delgerskhn/gosetup/pkg/domain/gorm"
	"github.com/delgerskhn/gosetup/pkg/infrastructure/persistence"
	repo "github.com/delgerskhn/gosetup/pkg/infrastructure/persistence/repository"
)

var repository *repo.GormRepository[gorm.BookGorm, domain.Book]

func init() {
	db := persistence.GetDB()
	repository = repo.NewRepository[gorm.BookGorm, domain.Book](db)
}

func ChangeAuthor(id string, author string) error {
	book, err := repository.FindByID(context.Background(), id)
	if err != nil {
		return err
	}
	updated := book.ChangeAuthor(author)
	return repository.Update(context.Background(), &updated)
}

func GetBooks() ([]domain.Book, error) {
	return repository.FindAll(context.Background())

}

func CreateBook(book domain.Book) (domain.Book, error) {
	err := repository.Insert(context.Background(), &book)
	return book, err
}

func GetBook(id string) (domain.Book, error) {
	return repository.FindByID(context.Background(), id)
}

func DeleteBook(id string) error {
	return repository.DeleteById(context.Background(), id)
}

func UpdateBook(book domain.Book) (domain.Book, error) {
	return book, repository.Update(context.Background(), &book)
}
