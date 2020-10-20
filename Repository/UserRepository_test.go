package Services_test

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/r-keegan/synoptic-project/Models"
	. "github.com/r-keegan/synoptic-project/Services"
	"testing"
)

var db *gorm.DB
var userService UserService

func TestUserService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Service Suite")
}

var _ = Describe("UserService", func() {

	BeforeSuite(func() {
		GetDatabase()
	})

	BeforeEach(func() {
		// setup database
		db.DropTableIfExists(&Models.User{})
		db.AutoMigrate(&Models.User{})

		// set up user service
		userService = UserService{DB: db}
	})

	AfterEach(func() {
		db.DropTableIfExists(&Models.User{})
	})

	Context("Create User", func() {
		It("creates a user when that user does not exist", func() {
			//Given I have a User
			user := getUserOne()

			//When I create that user
			err := userService.CreateUser(user)

			//then that user is created in the database
			Expect(err).ShouldNot(HaveOccurred())

			var foundUser Models.User
			db.Find(&foundUser)
			compare(foundUser, user)
		})

		It("creates multiple users", func() {
			//Given I have two users
			user1 := getUserOne()
			user2 := getUserTwo()

			//When I create both users
			err1 := userService.CreateUser(user1)
			err2 := userService.CreateUser(user2)

			//then that user is created in the database
			Expect(err1).To(BeNil())
			Expect(err2).To(BeNil())

			var foundUsers []Models.User
			db.Find(&foundUsers)

			compare(foundUsers[0], user1)
			compare(foundUsers[1], user2)
			Expect(foundUsers).To(HaveLen(2))
		})
	})
})

func GetDatabase() {
	var err error
	db, err = gorm.Open("sqlite3", "UserServiceTest.db")
	if err != nil {
		fmt.Println("Failed to connect to database: ", err)
	}
}

func getUserOne() Models.User {
	user := Models.User{
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
		EmployeeID: 4,
		Name:       "Maxeen Power",
		Email:      "maxeen.power@gmail.com",
		Phone:      "09716244907",
		Pin:        5432,
	}
	return user
}

//Test Helpers
//TODO investigate using assert.Contains(t, foundUsers, user1, user2) instead of two asserts that implicity check ordering
func compare(expectedUser Models.User, actualUser Models.User) {
	Expect(actualUser.EmployeeID).To(Equal(expectedUser.EmployeeID))
	Expect(actualUser.Name).To(Equal(expectedUser.Name))
	Expect(actualUser.Phone).To(Equal(expectedUser.Phone))
	Expect(actualUser.Pin).To(Equal(expectedUser.Pin))
}
