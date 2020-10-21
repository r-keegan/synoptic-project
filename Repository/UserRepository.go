package Repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/r-keegan/synoptic-project/Models"
)

type UserRepository struct {
	DB *gorm.DB
}

// TODO return user and an error
func (r UserRepository) CreateUser(user Models.User) (err error) {
	if err = r.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r UserRepository) GetUserByEmployeeID(employeeID int) (user Models.User, err error) {
	if err = r.DB.Where("employee_id = ?", employeeID).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r UserRepository) GetUserByCardID(cardID string) (user Models.User, err error) {
	if err = r.DB.Where("card_id = ?", cardID).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r UserRepository) UpdateUser(user Models.User) (err error) {
	r.DB.Save(user)
	fmt.Println(user)
	return nil
}

func (r UserRepository) DeleteUserByEmployeeID(employeeID int) (err error) {
	//we should look into the need for this user declaration, does it actually contain a user in it once its deleted
	var user Models.User
	//not actually getting the error if there was one to return
	//We might need a first query on here incase there are multiple users with the same employee id by magic to be safe
	r.DB.Where("employee_id = ?", employeeID).Delete(&user)
	return nil
}

func (r UserRepository) DeleteUserByCardID(cardID string) (err error) {
	var user Models.User
	//not actually getting the error if there was one to return
	r.DB.Where("card_id = ?", cardID).Delete(&user)
	return nil
}
