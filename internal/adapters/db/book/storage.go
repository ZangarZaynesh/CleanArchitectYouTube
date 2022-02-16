package book

import (
	"student/internal/domain/book"
	"student/internal/module"

	"go.mongodb.org/mongo-driver/mongo"
)

type bookStorage struct {
	db mongo.Database
}

func NewStorage(db mongo.Database) book.Storage {
	return &bookStorage{db: db}
}

func (bs *bookStorage) GetOne(uuid string) *module.Book {
	return nil
}

func (bs *bookStorage) GetAll(limit, offset int) []*module.Book {
	return nil
}

func (bs *bookStorage) Create(book *module.Book) *module.Book {
	return nil
}

func (bs *bookStorage) Delete(book *module.Book) error {
	return nil
}
