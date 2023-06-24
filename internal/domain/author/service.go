package author

import (
	"clean_architecture/internal/adapters/api/author"
	"context"
)

type Service interface {
	GetBookByUUID(ctx context.Context, uuid string) *Author
	GetAllBooks(ctx context.Context, limit, offset int) []*Author
	CreateBook(ctx context.Context, dto *CreateAuthorDto) *Author
}

type service struct {
	storage Storage
}

func NewService(storage Storage) author.Service {
	return &service{storage: storage}
}

func (s *service) Create(ctx context.Context, dto *CreateAuthorDto) *Author {
	return nil
}

func (s *service) GetByUUID(ctx context.Context, uuid string) *Author {
	return s.storage.GetOne(uuid)
}

func (s *service) GetAll(ctx context.Context, limit, offset int) []*Author {
	return s.storage.GetAll(limit, offset)
}
