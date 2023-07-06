package auth

import (
	"library/entity"

	"github.com/gin-gonic/gin"
)

type RepoAuth interface {
	Login(string) (*entity.Login, error)
	Regist(entity.Login) error
}

type UsecaseAuth interface {
	Login(ctx *gin.Context) error
	Regist(ctx *gin.Context) error
}
