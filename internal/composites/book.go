package composites

import (
	"clean_architecture/internal/adapters/api"
	book2 "clean_architecture/internal/adapters/api/book"
	book3 "clean_architecture/internal/adapters/db/book"
	"clean_architecture/internal/domain/book"
)

type BookComposite struct {
	Storage book.Storage
	Service book2.Service
	Handler api.Handler
}

func NewBookComposite() (*BookComposite, error) {
	storage := book3.NewStorage()
	service := book.NewService(storage)
	handler := book2.NewHandler(service)
	return &BookComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
