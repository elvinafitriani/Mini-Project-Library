package book

import (
	"library/entity"

	"github.com/gin-gonic/gin"
)

type UsecaseBook interface {
	CreateBook(*gin.Context) error
	GetAllBooks(ctx *gin.Context) ([]entity.Books, error)
	GetAuthorsByBook(ctx *gin.Context) ([]string, error)
	UpdateBook(ctx *gin.Context) error
	DeleteBook(ctx *gin.Context) error
}

type RepositoryBook interface {
	CreateBook(entity.Books, *gin.Context) error
	GetAllBooks() ([]entity.Books, error)
	GetAuthorsByBook(string) (*entity.Books, error)
	UpdateBook(entity.Books, string, *gin.Context) error
	DeleteBook(string) error
}
