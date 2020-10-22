package Repository

import (
	"github.com/jinzhu/gorm"
	"github.com/r-keegan/synoptic-project/Models"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r UserRepository) CreateUser(user Models.User) (err error) {
	if err = r.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r UserRepository) GetUserByCardID(cardID string) (user Models.User, err error) {
	if err = r.DB.Where("card_id = ?", cardID).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r UserRepository) UpdateUser(user Models.User) (err error) {
	r.DB.Save(user)
	return nil
}

func (r UserRepository) GetUserByEmployeeID(employeeID int) (user Models.User, err error) {
	if err = r.DB.Where("employee_id = ?", employeeID).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r UserRepository) DeleteUserByEmployeeID(employeeID int) (err error) {
	var user Models.User
	r.DB.Where("employee_id = ?", employeeID).Delete(&user)
	return nil
}

func (r UserRepository) DeleteUserByCardID(cardID string) (err error) {
	var user Models.User
	r.DB.Where("card_id = ?", cardID).Delete(&user)
	return nil
}
