package book

import (
	"clean_architecture/internal/domain/book"
	"context"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) *book.Book
	GetAll(ctx context.Context, limit, offset int) []*book.Book
	Create(ctx context.Context, dto *book.CreateBookDto) *book.Book
}
