package Services

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/r-keegan/synoptic-project/Models"
	"github.com/stretchr/testify/assert"
	"testing"
)

var db *gorm.DB
var userService UserService

func TestMain(m *testing.M) {
	// setup database
	GetDatabase23()
	db.DropTableIfExists(&Models.User{})
	db.AutoMigrate(&Models.User{})

	// set up user service
	userService = UserService{DB: db}

	// runs tests
	m.Run()
	db.DropTableIfExists(&Models.User{})
}

func GetDatabase23() {
	var err error
	db, err = gorm.Open("sqlite3", "UserServiceTest.db")
	if err != nil {
		fmt.Println("Failed to connect to database: ", err)
	}
}

func TestUserService_CreateUser(t *testing.T) {
	//Given I have a User
	user := getUserOne()

	//When I create that user
	err := userService.CreateUser(user)

	//then that user is created in the database
	assert.Nil(t, err)

	var foundUser Models.User
	db.Find(&foundUser)
	assert.Equal(t, foundUser, user)
}

func TestUserService_CreateMultipleUsers(t *testing.T) {
	//Given I have two users
	user1 := getUserOne()
	user2 := getUserTwo()

	//When I create both users
	err1 := userService.CreateUser(user1)
	err2 := userService.CreateUser(user2)

	//then that user is created in the database
	assert.Nil(t, err1)
	assert.Nil(t, err2)

	var foundUsers []Models.User
	db.Find(&foundUsers)

	assert.Equal(t, foundUsers, user1)
	assert.Equal(t, foundUsers[0], user1)
	assert.Equal(t, foundUsers[1], user2)
	assert.Equal(t, len(foundUsers), 2)
	//TODO investigate using assert.Contains(t, foundUsers, user1, user2) instead of two asserts that implicity check ordering
}

func getUserOne() Models.User {
	user := Models.User{
		ID:         1,
		EmployeeID: 2,
		Name:       "Max Power",
		Email:      "max.power@gmail.com",
		Phone:      "09716244907",
		Pin:        1234,
	}
	return user
}

func getUserTwo() Models.User {
	user := Models.User{
		ID:         2,
		EmployeeID: 4,
		Name:       "Maxeen Power",
		Email:      "maxeen.power@gmail.com",
		Phone:      "09716244907",
		Pin:        5432,
	}
	return user
}
