package author

import (
	"library/entity"

	"github.com/gin-gonic/gin"
)

type UsecaseAuthor interface {
	CreateAuthor(*gin.Context) error
	GetAllAuthors(ctx *gin.Context) ([]entity.Authors, error)
	GetBooksByAuthor(ctx *gin.Context) ([]string, error)
	UpdateAuthor(ctx *gin.Context) error
	DeleteAuthor(ctx *gin.Context) error
}

type RepositoryAuthor interface {
	CreateAuthor(entity.Authors, *gin.Context) error
	GetAllAuthors() ([]entity.Authors, error)
	GetBooksByAuthor(string) (*entity.Authors, error)
	UpdateAuthor(entity.Authors, string, *gin.Context) error
	DeleteAuthor(string) error
}
