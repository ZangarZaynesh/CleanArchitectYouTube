package author

import (
	"context"
	"student/internal/module"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) *module.Author
	GetAll(ctx context.Context, limit, offset int) []*module.Author
	Create(ctx context.Context, dto *module.CreateAuthorDTO) *module.Author
}
