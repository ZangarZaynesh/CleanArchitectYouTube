package module

import (
	"fmt"
)

type Author struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Book struct {
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Author Author `json:"author"`
	Busy   bool   `json:"busy"`
	Owner  string `json:"owner"`
}

func (b *Book) Take(owner string) error {
	if b.Busy {
		return fmt.Errorf("book is busy")
	}
	b.Owner = owner
	b.Busy = true
	return nil
}
