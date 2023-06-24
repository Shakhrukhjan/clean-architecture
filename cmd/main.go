package main

import (
	author3 "clean_architecture/internal/adapters/api/author"
	author2 "clean_architecture/internal/adapters/db/author"
	"clean_architecture/internal/composites"
	"clean_architecture/internal/domain/author"
)

func main() {
	// entry point
	authorComposite := composites.NewAuthorComposite()
	authorStorage := author2.NewStorage()
	authorService := author.NewService(authorStorage)
	authorHandler := author3.NewHandler(authorService)

}
