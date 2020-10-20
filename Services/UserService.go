package Services

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/r-keegan/synoptic-project/Models"
	"github.com/r-keegan/synoptic-project/Repository"
	"strings"
)

type UserRepository interface {
	CreateUser(Models.User) error
}

type UserService struct {
	UserRepository Repository.UserRepository
}

func (s UserService) CreateUser(user Models.User) (err error) {
	err = s.Validate(user, "update")
	if err != nil {
		return err
	}
	err = s.UserRepository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
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

func (s UserService) Validate(user Models.User, action string) (err error) {
	switch strings.ToLower(action) {
	case "update":
		if user.EmployeeID < 1 {
			return errors.New("Required employeeID")
		}
		if user.CardID == "" {
			return errors.New("Required cardID")
		}
		if !(len(user.CardID) == 16) {
			return errors.New("Invalid cardID")
		}
		if user.Name == "" {
			return errors.New("Required name")
		}
		if user.Email == "" {
			return errors.New("Required email")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid email")
		}
		if user.Phone == "" {
			return errors.New("Required phone")
		}
	}
	return nil
}
