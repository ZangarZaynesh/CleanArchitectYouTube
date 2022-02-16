package author

import (
	"context"
	"student/internal/adapters/api/author"
	"student/internal/module"
)

type service struct {
	storage Storage
}

func NewService(storage Storage) author.Service {
	return &service{storage: storage}
}

func (s *service) Create(ctx context.Context, dto *module.CreateAuthorDTO) *module.Author {
	return nil
}

func (s *service) GetByUUID(ctx context.Context, uuid string) *module.Author {
	return s.storage.GetOne(uuid)
}

func (s *service) GetAll(ctx context.Context, limit, offset int) []*module.Author {
	return s.storage.GetAll(limit, offset)
}
