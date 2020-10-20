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
		if !(user.Pin == user.ConfirmPin) {
			return errors.New("Pin does not match Confirmation Pin")
		}
	}
	return nil
}
