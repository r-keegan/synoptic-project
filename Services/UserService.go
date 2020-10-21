package Services

import (
	"errors"
	"fmt"
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

func (s UserService) Validate(user Models.User, action string) (err error) {
	//TODO perhaps throw a different exception
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

func (s UserService) Authenticate(userAuth Models.AuthenticatedRequest) bool {
	_, err := s.getAUserByCardAndPin(userAuth.CardID, userAuth.Pin)
	if err == nil {
		return true
	}
	return false
}

func (s UserService) GetBalance(cardID string, pin string) (int, error) {
	user, err := s.getAUserByCardAndPin(cardID, pin)
	if err == nil {
		return user.Balance, nil
	}
	return 0, err
}

func (s UserService) Purchase(cardID string, pin string, amount int) (int, error) {
	user, err := s.getAUserByCardAndPin(cardID, pin)
	if err == nil {
		potentialBalance := user.Balance - amount
		if potentialBalance > 0 {
			user.Balance = potentialBalance
			err = s.UpdateUser(user)
			if err == nil {
				return user.Balance, nil
			}
		}
	}
	return user.Balance, errors.New(fmt.Sprintf("Unable to make purchase: your balance is %v", user.Balance))
}

func (s UserService) TopUp(cardID string, pin string, amount int) (int, error) {
	user, err := s.getAUserByCardAndPin(cardID, pin)
	if err == nil {
		user.Balance = user.Balance + amount
		err = s.UpdateUser(user)
		if err == nil {
			return user.Balance, nil
		}
	}
	return user.Balance, errors.New(fmt.Sprintf("Unable to topup: your balance is %v", user.Balance))
}

//todo make private
func (s UserService) getAUserByCardAndPin(cardID string, pin string) (Models.User, error) {
	user, err := s.GetEmployeeByCardID(cardID)
	if err == nil {
		if user.Pin == pin {
			return user, nil
		}
	}
	return user, err
}

func (s UserService) CreateUser(createUser Models.CreateUser) error {
	user := Models.User{
		EmployeeID: createUser.EmployeeID,
		CardID:     createUser.CardID,
		Name:       createUser.Name,
		Email:      createUser.Email,
		Phone:      createUser.Phone,
		Pin:        createUser.Pin,
		Balance:    0,
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
