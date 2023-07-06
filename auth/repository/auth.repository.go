package repository

import (
	"library/entity"

	"gorm.io/gorm"
)

func NewRepoAuth(db *gorm.DB) Database {
	return Database{
		Db: db,
	}
}

type Database struct {
	Db *gorm.DB
}

func (db Database) Login(username string) (*entity.Login, error) {
	var user entity.Login

	if err := db.Db.First(&user, "username=?", username).Error; err != nil {
		return nil, err
	}

	err := db.Db.Find(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db Database) Regist(user entity.Login) error {
	err := db.Db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}
