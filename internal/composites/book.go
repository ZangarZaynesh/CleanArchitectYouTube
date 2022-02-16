package composites

import (
	api "student/internal/adapters/api"
	book2 "student/internal/adapters/api/book"
	book3 "student/internal/adapters/db/book"
	"student/internal/domain/book"
)

type BookComposite struct {
	Storage book.Storage
	Service book2.Service
	Handler api.Handler
}

func NewBookComposite(mongoComposite *MongoDBComposite) (*BookComposite, error) {
	storage := book3.NewStorage(*mongoComposite.db)
	service := book.NewService(storage)
	handler := book2.NewHandler(service)
	return &BookComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
