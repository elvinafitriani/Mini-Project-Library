package handlers

import (
	"library/author"
	"library/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandlersAuthor(usecs author.UsecaseAuthor, r *gin.RouterGroup) {
	eng := &usecase{
		use: usecs,
	}

	v2 := r.Group("author")
	v2.POST("", middleware.Auth(), eng.CreateAuthor)
	v2.GET("", middleware.Auth(), eng.GetAllAuthors)
	v2.GET("/:name", middleware.Auth(), eng.GetBooksByAuthor)
	v2.PUT("/:name", middleware.Auth(), eng.UpdateAuthor)
	v2.DELETE("/:name", middleware.Auth(), eng.DeleteAuthor)

}

type usecase struct {
	use author.UsecaseAuthor
}

func (author usecase) CreateAuthor(ctx *gin.Context) {
	err := author.use.CreateAuthor(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Data created successfully."})
}

func (author usecase) GetAllAuthors(ctx *gin.Context) {
	result, err := author.use.GetAllAuthors(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Authors": result, "Response": "Data retrieved successfully."})
}

func (author usecase) GetBooksByAuthor(ctx *gin.Context) {
	result, err := author.use.GetBooksByAuthor(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (author usecase) UpdateAuthor(ctx *gin.Context) {
	err := author.use.UpdateAuthor(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Data updated successfully."})
}

func (author usecase) DeleteAuthor(ctx *gin.Context) {
	err := author.use.DeleteAuthor(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Response": "Data deleted successfully."})
}
