package Models

import (
	"github.com/jinzhu/gorm"
	"github.com/r-keegan/synoptic-project/Config"
)

func GetDatabase() *gorm.DB {
	return Config.DB
}

func GetAllUsers(user *[]User) (err error) {
	db := GetDatabase()
	find := db.Find(user)
	if err = find.Error; err != nil {
		return err
	}
	return nil
}
