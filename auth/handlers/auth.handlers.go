package handlers

import (
	"library/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandlersAuth(us auth.UsecaseAuth, r *gin.RouterGroup) {
	eng := &Usecase{
		usAuth: us,
	}

	v2 := r.Group("auth")
	v2.POST("/login", eng.Login)
	v2.POST("/regist", eng.Regist)
}

type Usecase struct {
	usAuth auth.UsecaseAuth
}

func (us Usecase) Login(ctx *gin.Context) {
	err := us.usAuth.Login(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Token JWT disimpan dalam cookie 'token'"})
}

func (us Usecase) Regist(ctx *gin.Context) {
	err := us.usAuth.Regist(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Success Regist"})
}
