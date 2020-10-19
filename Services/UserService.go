package Services

import (
	"github.com/jinzhu/gorm"
	"github.com/r-keegan/synoptic-project/Models"
)

type UserService struct {
	DB *gorm.DB
}

//func GetAllUsers(user *[]Models.User) (err error) {
//	db := GetDatabase()
//	find := db.Find(user)
//	if err = find.Error; err != nil {
//		return err
//	}
//	return nil
//}

func (s UserService) CreateUser(user Models.User) (err error) {
	if err = s.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//func GetUserByID(user *Models.User, id string) (err error) {
//	db := GetDatabase()
//	if err = db.Where("id = ?", id).First(user).Error; err != nil {
//		return err
//	}
//	return nil
//}
