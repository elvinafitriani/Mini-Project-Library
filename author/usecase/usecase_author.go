package usecase

import (
	"encoding/json"
	"errors"
	"library/author"
	"library/entity"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pariz/gountries"
)

func NewUsecaseAuthor(repoAuthor author.RepositoryAuthor) repositoryAuthor {
	return repositoryAuthor{
		repo: repoAuthor,
	}
}

type repositoryAuthor struct {
	repo author.RepositoryAuthor
}

func (author repositoryAuthor) CreateAuthor(ctx *gin.Context) error {
	var authors entity.Authors

	if err := ctx.ShouldBindJSON(&authors); err != nil {
		return err
	}

	if len(authors.Name) < 3 || len(authors.Name) > 20 {
		err := errors.New("panjang nama tidak sesuai")
		return err
	}

	invalidChars := []string{"!", "@", "#", "$", "%"}
	for _, char := range invalidChars {
		if strings.Contains(authors.Name, char) {
			err := errors.New("nama mengandung karakter yang tidak diizinkan")
			return err
		}
	}

	data := gountries.New()

	_, err := data.FindCountryByName(authors.Country)
	if err != nil {
		err := errors.New("invalid country")
		return err
	}
	serializedBooks, err := json.Marshal(authors.Book)
	if err != nil {
		return err
	}

	authors.SerializedBooks = string(serializedBooks)

	if err := author.repo.CreateAuthor(authors, ctx); err != nil {
		return err
	}

	return nil
}

func (author repositoryAuthor) GetAllAuthors(ctx *gin.Context) ([]entity.Authors, error) {
	var authors []entity.Authors
	var err error

	authors, err = author.repo.GetAllAuthors()
	if err != nil {
		return nil, err
	}

	for i, v := range authors {
		var book []string
		err := json.Unmarshal([]byte(v.SerializedBooks), &book)
		if err != nil {
			return nil, err
		}
		authors[i].Book = book
	}

	return authors, nil
}

func (author repositoryAuthor) GetBooksByAuthor(ctx *gin.Context) ([]string, error) {
	var Author struct {
		Name string `uri:"name"`
	}

	if err := ctx.ShouldBindUri(&Author); err != nil {
		return nil, err
	}

	books, err := author.repo.GetBooksByAuthor(Author.Name)
	if err != nil {
		return nil, err
	}

	var book []string
	if err := json.Unmarshal([]byte(books.SerializedBooks), &book); err != nil {
		return nil, err
	}
	books.Book = book

	return books.Book, nil
}

func (author repositoryAuthor) UpdateAuthor(ctx *gin.Context) error {
	var authors entity.Authors
	type Authors struct {
		Name    string   `json:"name" binding:"required"`
		Country string   `json:"country" binding:"required"`
		Book    []string ` json:"book" binding:"required"`
	}

	var Name struct {
		Name string `uri:"name"`
	}

	var authorreq Authors

	if err := ctx.ShouldBindJSON(&authorreq); err != nil {
		return err
	}

	if err := ctx.ShouldBindUri(&Name); err != nil {
		return err
	}

	authors.Name = authorreq.Name
	authors.Country = authorreq.Country
	authors.Book = authorreq.Book

	if len(authors.Name) < 3 || len(authors.Name) > 20 {
		err := errors.New("panjang nama tidak sesuai")
		return err
	}

	invalidChars := []string{"!", "@", "#", "$", "%"}
	for _, char := range invalidChars {
		if strings.Contains(authors.Name, char) {
			err := errors.New("nama mengandung karakter yang tidak diizinkan")
			return err
		}
	}

	data := gountries.New()

	_, err := data.FindCountryByName(authors.Country)
	if err != nil {
		err := errors.New("invalid country")
		return err
	}

	serializedBooks, err := json.Marshal(authors.Book)
	if err != nil {
		return err
	}

	authors.SerializedBooks = string(serializedBooks)

	if err := author.repo.UpdateAuthor(authors, Name.Name, ctx); err != nil {
		return err
	}

	return nil
}

func (author repositoryAuthor) DeleteAuthor(ctx *gin.Context) error {
	var Name struct {
		Name string `uri:"name"`
	}

	if err := ctx.ShouldBindUri(&Name); err != nil {
		return err
	}

	err := author.repo.DeleteAuthor(Name.Name)

	if err != nil {
		return err
	}

	return nil
}
