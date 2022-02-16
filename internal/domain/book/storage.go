package book

import "student/internal/module"

type Storage interface {
	GetOne(uuid string) *module.Book
	GetAll(limit, offset int) []*module.Book
	Create(book *module.Book) *module.Book
	Delete(book *module.Book) error
}
