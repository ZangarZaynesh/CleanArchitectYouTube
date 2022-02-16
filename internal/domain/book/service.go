package book

import (
	"context"
	"student/internal/adapters/api/author"
	"student/internal/adapters/api/book"
	"student/internal/module"
)

type service struct {
	storage       Storage
	authorService author.Service
}

func NewService(storage Storage) book.Service {
	return &service{storage: storage}
}

func (s *service) GetByUUID(ctx context.Context, uuid string) *module.Book {
	return s.storage.GetOne(uuid)
}

func (s *service) GetAll(ctx context.Context, limit, offset int) []*module.Book {
	return s.storage.GetAll(limit, offset)
}

func (s *service) Create(ctx context.Context, dto *module.CreateBookDTO) *module.Book {
	author := s.authorService.GetByUUID(ctx, dto.AuthorUUID)
	return nil
}
