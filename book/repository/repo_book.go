package repository

import (
	"errors"
	"library/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRepoBook(data *gorm.DB) Database {
	return Database{
		db: data,
	}
}

type Database struct {
	db *gorm.DB
}

func (db Database) CreateBook(book entity.Books, ctx *gin.Context) error {
	var author *entity.Authors
	for i, v := range book.Author {
		err := db.db.Find(&author, "name=?", v).Error
		if err != nil {
			return err
		}
		if i != 0 {
			if v == book.Author[i-1] {
				err = errors.New("author duplicate")
				return err
			}
			if v == "" {
				err = errors.New("author nil")
				return err
			}
		}
	}

	err := db.db.Where("isbn = ?", book.ISBN).First(&book)
	if err != nil {
		if err := db.db.Create(&book).Error; err != nil {
			return err
		}
	} else {
		err := errors.New("isbn sudah terdaftar")
		return err
	}

	return nil
}

func (db Database) GetAllBooks() (book []entity.Books, err error) {
	err = db.db.Find(&book).Error
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (db Database) GetAuthorsByBook(isbn string) (*entity.Books, error) {
	var book *entity.Books

	if err := db.db.First(&book, "isbn=?", isbn).Error; err != nil {
		return nil, err
	}

	err := db.db.Find(&book, "isbn=?", isbn).Error
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (db Database) UpdateBook(Book entity.Books, isbn string, ctx *gin.Context) error {
	var author *entity.Authors
	for i, v := range Book.Author {
		err := db.db.Find(&author, "name=?", v).Error
		if err != nil {
			return err
		}
		if i != 0 {
			if v == Book.Author[i-1] {
				err = errors.New("author duplicate")
				return err
			}
			if v == "" {
				err = errors.New("author nil")
				return err
			}
		}
	}

	if err := db.db.First(&entity.Books{}, "isbn=?", isbn).Error; err != nil {
		return err
	}
	if err := db.db.Where("isbn=?", isbn).Updates(&Book).Error; err != nil {
		return err
	}
	return nil
}

func (db Database) DeleteBook(isbn string) error {
	if err := db.db.First(&entity.Books{}, "isbn=?", isbn).Error; err != nil {
		return err
	}

	if err := db.db.Where("isbn=?", isbn).Delete(&entity.Books{}).Error; err != nil {
		return err
	}

	return nil
}
