package composites

import (
	api "student/internal/adapters/api"
	"student/internal/adapters/api/author"
	author3 "student/internal/adapters/db/author"
	author2 "student/internal/domain/author"
)

type AuthorComposite struct {
	Storage author2.Storage
	Service author.Service
	Handler api.Handler
}

func NewAuthorComposite(mongoComposite *MongoDBComposite) (*AuthorComposite, error) {
	storage := author3.NewStorage(*mongoComposite.db)
	service := author2.NewService(storage)
	handler := author.NewHandler(service)
	return &AuthorComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
