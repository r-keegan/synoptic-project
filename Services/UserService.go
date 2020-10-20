package Services

import (
	"github.com/r-keegan/synoptic-project/Models"
	"github.com/r-keegan/synoptic-project/Repository"
)

type UserService struct {
	UserRepository Repository.UserRepository
}

// add validation
func (s UserService) CreateUser(user Models.User) (err error) {
	return s.UserRepository.CreateUser(user)
}

//func (s UserService) UpdateCardBalance(userId int, newBalance int) {
//
//}

//func (s UserService) GetUserByID(userId string) (err error) {
//	if err = s.DB.Where("id = ?", id).First(user).Error; err != nil {
//		return err
//	}
//	return nil
//}
