package Services

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

func GetUserByID(user *Models.User, id string) (err error) {
	db := GetDatabase()
	if err = db.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}
