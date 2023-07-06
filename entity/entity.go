package entity

import "gorm.io/gorm"

type Login struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" binding:"min=8"`
}

type Books struct {
	gorm.Model
	Title             string    `json:"title" binding:"required"`
	PublishedYear     int       `json:"publishedYear" binding:"required" `
	ISBN              string    `json:"isbn" binding:"required" gorm:"unique"`
	Author            []string  `json:"author" gorm:"-"`
	SerializedAuthors string    `gorm:"column:author" json:"-"`
	Authors           []Authors `gorm:"many2many:book_authors;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

type Authors struct {
	gorm.Model
	Name            string   `json:"name" binding:"required" gorm:"unique"`
	Country         string   `json:"country" binding:"required"`
	Book            []string `json:"book" gorm:"-"`
	SerializedBooks string   `gorm:"column:book" json:"-"`
	Books           []Books  `gorm:"many2many:book_authors;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
