package author

import "student/internal/module"

type Storage interface {
	GetOne(uuid string) *module.Author
	GetAll(limit, offset int) []*module.Author
	Create(book *module.Author) *module.Author
	Delete(book *module.Author) error
}
