package app

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

type Repository struct {
	db *gorm.DB
}

func NewRepository() Repository {
	// gorm
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database!~!")
	}

	// create table
	db.AutoMigrate(&User{})

	return Repository{db}
}

func (r Repository) createUser(username string, password string) (int64, error) {
	res := r.db.Create(&User{
		Username: username,
		Password: password,
	})

	return res.RowsAffected, res.Error
}

func (r Repository) findUser(username string) (User, error) {
	var user User
	res := r.db.First(&user, "username = ?", username)

	return user, res.Error
}
