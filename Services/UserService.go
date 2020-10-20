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

func (s UserService) UpdateUser(user Models.User) (err error) {
	err = s.Validate(user, "update")
	if err != nil {
		return err
	}
	err = s.UserRepository.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

//func (s UserService) UpdateCardBalance(userId int, newBalance int) {
//
//}

func (s UserService) Validate(user Models.User, action string) (err error) {
	//TODO perhaps throw a different exeption
	switch strings.ToLower(action) {
	case "update":
		if user.EmployeeID < 1 {
			return errors.New("Required employeeID")
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
		if !(len(user.Pin) == 4) {
			return errors.New("Invalid pin")
		}
	}
	return nil
}

func (s UserService) GetEmployeeByCardID(cardID string) (Models.User, error) {
	return s.UserRepository.GetUserByCardID(cardID)
}

func (s UserService) Authenticate(userAuth Models.UserAuth) bool {
	user, err := s.GetEmployeeByCardID(userAuth.CardID)
	if err == nil {
		if user.Pin == userAuth.Pin {
			return true
		}
	}
	return false
}

func (s UserService) CreateUser(createUser Models.CreateUser) error {
	user := Models.User{
		EmployeeID: createUser.EmployeeID,
		CardID:     createUser.CardID,
		Name:       createUser.Name,
		Email:      createUser.Phone,
		Phone:      createUser.Phone,
		Pin:        createUser.Pin,
	}
	return s.CreateUser2(user)
}

func (s UserService) CreateUser2(user Models.User) (err error) {
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
