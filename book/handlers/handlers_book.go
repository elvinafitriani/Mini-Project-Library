package handlers

import (
	"library/book"
	"library/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandlersBook(usecs book.UsecaseBook, r *gin.RouterGroup) {
	eng := &usecase{
		use: usecs,
	}

	v2 := r.Group("book")
	v2.POST("/", middleware.Auth(), eng.CreateBook)
	v2.GET("", middleware.Auth(), eng.GetAllBooks)
	v2.GET("/:isbn", middleware.Auth(), eng.GetAuthorsByBook)
	v2.PUT("/:isbn", middleware.Auth(), eng.UpdateBook)
	v2.DELETE("/:isbn", middleware.Auth(), eng.DeleteBook)

}

type usecase struct {
	use book.UsecaseBook
}

func (book usecase) CreateBook(ctx *gin.Context) {
	err := book.use.CreateBook(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Data created successfully."})
}

func (book usecase) GetAllBooks(ctx *gin.Context) {
	result, err := book.use.GetAllBooks(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Books": result, "Response": "Data retrieved successfully."})
}

func (book usecase) GetAuthorsByBook(ctx *gin.Context) {
	result, err := book.use.GetAuthorsByBook(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (book usecase) UpdateBook(ctx *gin.Context) {
	err := book.use.UpdateBook(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Data updated successfully."})
}

func (book usecase) DeleteBook(ctx *gin.Context) {
	err := book.use.DeleteBook(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Response": "Data deleted successfully."})
}
