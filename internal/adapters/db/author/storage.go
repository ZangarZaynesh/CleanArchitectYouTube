package author

import (
	"student/internal/domain/author"
	"student/internal/module"

	"go.mongodb.org/mongo-driver/mongo"
)

type authorStorage struct {
	db mongo.Database
}

func NewStorage(db mongo.Database) author.Storage {
	return &authorStorage{db: db}
}

func (bs *authorStorage) GetOne(uuid string) *module.Author {
	return nil
}

func (bs *authorStorage) GetAll(limit, offset int) []*module.Author {
	return nil
}

func (bs *authorStorage) Create(book *module.Author) *module.Author {
	return nil
}

func (bs *authorStorage) Delete(book *module.Author) error {
	return nil
}
