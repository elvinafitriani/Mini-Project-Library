package usecase

import (
	"library/auth"
	"library/entity"
	"library/jwt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func NewUsecaseAuth(rep auth.RepoAuth) Repo {
	return Repo{
		repAuth: rep,
	}
}

type Repo struct {
	repAuth auth.RepoAuth
}

func (rep Repo) Login(ctx *gin.Context) error {
	var user entity.Login
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return err
	}

	result, err := rep.repAuth.Login(user.Username)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
		return err
	}

	token, err := jwt.Token(int(result.ID), result.Username)
	if err != nil {
		return err
	}

	ctx.SetCookie("token", token.AccsesToken, 3000, "/", "localhost", false, true)
	return nil
}

func (rep Repo) Regist(ctx *gin.Context) error {
	var user entity.Login

	if err := ctx.ShouldBindJSON(&user); err != nil {
		return err
	}

	hassPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}

	user.Password = string(hassPass)

	err = rep.repAuth.Regist(user)
	if err != nil {
		return err
	}

	return nil
}