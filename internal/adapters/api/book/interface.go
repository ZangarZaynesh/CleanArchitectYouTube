package book

import (
	"context"
	"student/internal/module"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) *module.Book
	GetAll(ctx context.Context, limit, offset int) []*module.Book
	Create(ctx context.Context, dto *module.CreateBookDTO) *module.Book
}
