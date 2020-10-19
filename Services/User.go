package Services

import (
	"github.com/jinzhu/gorm"
	"github.com/r-keegan/synoptic-project/Config"
	"github.com/r-keegan/synoptic-project/Models"
)

func GetDatabase() *gorm.DB {
	return Config.DB
}

func GetAllUsers(user *[]Models.User) (err error) {
	db := GetDatabase()
	find := db.Find(user)
	if err = find.Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *Models.User) (err error) {
	db := GetDatabase()
	if err = db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
