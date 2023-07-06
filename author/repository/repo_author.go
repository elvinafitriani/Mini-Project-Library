package repository

import (
	"errors"
	"library/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRepoAuthor(data *gorm.DB) Database {
	return Database{
		db: data,
	}
}

type Database struct {
	db *gorm.DB
}

func (db Database) CreateAuthor(author entity.Authors, ctx *gin.Context) error {
	var book *entity.Books
	for i, v := range author.Book {
		err := db.db.Find(&book, "isbn=?", v).Error
		if err != nil {
			return err
		}
		if i != 0 {
			if v == author.Book[i-1] {
				err = errors.New("book duplicate")
				return err
			}
			if v == "" {
				err = errors.New("book nil")
				return err
			}
		}
	}

	if err := db.db.Create(&author).Error; err != nil {
		return err
	}
	return nil
}

func (db Database) GetAllAuthors() (author []entity.Authors, err error) {
	err = db.db.Find(&author).Error
	if err != nil {
		return nil, err
	}

	return author, nil
}

func (db Database) GetBooksByAuthor(name string) (*entity.Authors, error) {
	var author *entity.Authors
	if err := db.db.First(&author, "name=?", name).Error; err != nil {
		return nil, err
	}

	err := db.db.Find(&author, "name=?", name).Error
	if err != nil {
		return nil, err
	}

	return author, nil
}

func (db Database) UpdateAuthor(Author entity.Authors, name string, ctx *gin.Context) error {
	var book *entity.Books
	for i, v := range Author.Book {
		err := db.db.Find(&book, "isbn=?", v).Error
		if err != nil {
			return err
		}
		if i != 0 {
			if v == Author.Book[i-1] {
				err = errors.New("book duplicate")
				return err
			}
			if v == "" {
				err = errors.New("book nil")
				return err
			}
		}
	}
	if err := db.db.First(&entity.Authors{}, "name=?", name).Error; err != nil {
		return err
	}
	if err := db.db.Where("name=?", name).Updates(&Author).Error; err != nil {
		return err
	}
	return nil
}

func (db Database) DeleteAuthor(name string) error {
	if err := db.db.First(&entity.Authors{}, "name=?", name).Error; err != nil {
		return err
	}

	if err := db.db.Where("name=?", name).Delete(&entity.Authors{}).Error; err != nil {
		return err
	}

	return nil
}
