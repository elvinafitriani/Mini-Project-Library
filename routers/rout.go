package routers

import (
	handAuth "library/auth/handlers"
	repoAuth "library/auth/repository"
	useAuth "library/auth/usecase"

	handBook "library/book/handlers"
	repoBook "library/book/repository"
	useBook "library/book/usecase"

	handAuthor "library/author/handlers"
	repoAuthor "library/author/repository"
	useAuthor "library/author/usecase"

	"library/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routes struct {
	Db *gorm.DB
	R  *gin.Engine
}

func (r Routes) Routers() {
	middleware.Add(r.R, middleware.CORSMiddleware())
	v1 := r.R.Group("library")

	//auth
	repositoryAuth := repoAuth.NewRepoAuth(r.Db)
	usecaseAuth := useAuth.NewUsecaseAuth(repositoryAuth)
	handAuth.NewHandlersAuth(usecaseAuth, v1)

	//book
	repositoryBook := repoBook.NewRepoBook(r.Db)
	usecaseBook := useBook.NewUsecaseBook(repositoryBook)
	handBook.NewHandlersBook(usecaseBook, v1)

	//author
	repositoryAuthor := repoAuthor.NewRepoAuthor(r.Db)
	usecaseAuthor := useAuthor.NewUsecaseAuthor(repositoryAuthor)
	handAuthor.NewHandlersAuthor(usecaseAuthor, v1)
}
