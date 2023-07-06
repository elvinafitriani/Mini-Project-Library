package usecase

import (
	"encoding/json"
	"errors"
	"library/book"
	"library/entity"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type getBooks struct {
	Title         string
	PublishedYear int
	ISBN          string
	Author        []string
}

func NewUsecaseBook(repoBook book.RepositoryBook) repository {
	return repository{
		repo: repoBook,
	}
}

type repository struct {
	repo book.RepositoryBook
}

func (book repository) CreateBook(ctx *gin.Context) error {
	var books entity.Books
	if err := ctx.ShouldBindJSON(&books); err != nil {
		return err
	}

	currentYear := time.Now().Year()
	if books.PublishedYear > currentYear {
		err := errors.New("tahun terbit tidak boleh melebihi tahun sekarang")
		return err
	}
	if books.PublishedYear < 1900 {
		err := errors.New("tahun terbit harus lebih besar dari 1900")
		return err
	}

	books.ISBN = strings.ReplaceAll(books.ISBN, "-", "")

	if len(books.ISBN) != 10 && len(books.ISBN) != 13 {
		err := errors.New("Nomor ISBN harus terdiri dari 10 atau 13 digit")
		return err
	}

	match, _ := regexp.MatchString(`^\d+$`, books.ISBN)
	if !match {
		err := errors.New("Nomor ISBN hanya boleh berisi digit angka")
		return err
	}

	serializedAuthors, err := json.Marshal(books.Author)
	if err != nil {
		return err
	}

	books.SerializedAuthors = string(serializedAuthors)

	
	if err := book.repo.CreateBook(books, ctx); err != nil {
		return err
	}

	return nil
}

func (book repository) GetAllBooks(ctx *gin.Context) ([]entity.Books, error) {
	var books []entity.Books
	var err error

	books, err = book.repo.GetAllBooks()
	if err != nil {
		return nil, err
	}

	for i, v := range books {
		var author []string
		err := json.Unmarshal([]byte(v.SerializedAuthors), &author)
		if err != nil {
			return nil, err
		}
		books[i].Author = author
	}

	return books, nil
}

func (book repository) GetAuthorsByBook(ctx *gin.Context) ([]string, error) {
	var Book struct {
		Isbn string `uri:"isbn"`
	}

	if err := ctx.ShouldBindUri(&Book); err != nil {
		return nil, err
	}

	authors, err := book.repo.GetAuthorsByBook(Book.Isbn)
	if err != nil {
		return nil, err
	}

	var author []string
	if err := json.Unmarshal([]byte(authors.SerializedAuthors), &author); err != nil {
		return nil, err
	}
	authors.Author = author

	return authors.Author, nil
}

func (book repository) UpdateBook(ctx *gin.Context) error {
	var books entity.Books
	type Books struct {
		Title         string   `json:"title"  binding:"required"`
		PublishedYear int      `json:"publishedYear" binding:"required"`
		ISBN          string   `json:"isbn" binding:"required"`
		Author        []string `json:"author" binding:"required"`
	}

	var ISBN struct {
		Isbn string `uri:"isbn"`
	}

	var bookreq Books

	if err := ctx.ShouldBindJSON(&bookreq); err != nil {
		return err
	}

	if err := ctx.ShouldBindUri(&ISBN); err != nil {
		return err
	}

	books.Title = bookreq.Title
	books.PublishedYear = bookreq.PublishedYear
	books.ISBN = bookreq.ISBN
	books.Author = bookreq.Author

	currentYear := time.Now().Year()
	if books.PublishedYear > currentYear {
		err := errors.New("tahun terbit tidak boleh melebihi tahun sekarang")
		return err
	}
	if books.PublishedYear < 1900 {
		err := errors.New("tahun terbit harus lebih besar dari 1900")
		return err
	}

	books.ISBN = strings.ReplaceAll(books.ISBN, "-", "")

	if len(books.ISBN) != 10 && len(books.ISBN) != 13 {
		err := errors.New("Nomor ISBN harus terdiri dari 10 atau 13 digit")
		return err
	}

	match, _ := regexp.MatchString(`^\d+$`, books.ISBN)
	if !match {
		err := errors.New("Nomor ISBN hanya boleh berisi digit angka")
		return err
	}

	serializedAuthors, err := json.Marshal(books.Author)
	if err != nil {
		return err
	}

	books.SerializedAuthors = string(serializedAuthors)

	if err := book.repo.UpdateBook(books, ISBN.Isbn, ctx); err != nil {
		return err
	}

	return nil
}

func (book repository) DeleteBook(ctx *gin.Context) error {
	var ISBN struct {
		Isbn string `uri:"isbn"`
	}

	if err := ctx.ShouldBindUri(&ISBN); err != nil {
		return err
	}

	err := book.repo.DeleteBook(ISBN.Isbn)

	if err != nil {
		return err
	}

	return nil
}
