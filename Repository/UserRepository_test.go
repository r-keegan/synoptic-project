package Repository_test

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/r-keegan/synoptic-project/Models"
	. "github.com/r-keegan/synoptic-project/Repository"
	"testing"
)

var db *gorm.DB
var userRepository UserRepository

func TestUserRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Service Suite")
}

var _ = Describe("UserRepository", func() {

	BeforeSuite(func() {
		GetDatabase()
	})

	BeforeEach(func() {
		// setup database
		db.DropTableIfExists(&Models.User{})
		db.AutoMigrate(&Models.User{})

		// set up user service
		userRepository = UserRepository{DB: db}
	})

	AfterEach(func() {
		db.DropTableIfExists(&Models.User{})
	})

	Context("Create User", func() {
		It("creates a user when that user does not exist", func() {
			//Given I have a User
			user := getUserOne()

			//When I create that user
			err := userRepository.CreateUser(user)

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
			err1 := userRepository.CreateUser(user1)
			err2 := userRepository.CreateUser(user2)

			//then that user is created in the database
			Expect(err1).To(BeNil())
			Expect(err2).To(BeNil())

			var foundUsers []Models.User
			db.Find(&foundUsers)

			compare(foundUsers[0], user1)
			compare(foundUsers[1], user2)
			Expect(foundUsers).To(HaveLen(2))
		})

		It("throws an error when it employeeID already exists", func() {
			user1 := getUserOne()
			user2 := Models.User{
				EmployeeID: 2,
				CardID:     "b51TG7dqBy5wGO4L",
				Name:       "Max Power",
				Email:      "max.power@gmail.com",
				Phone:      "09716244907",
				Pin:        "1234",
				ConfirmPin: "1234",
			}
			err1 := userRepository.CreateUser(user1)
			err2 := userRepository.CreateUser(user2)

			Expect(err1).To(BeNil())
			Expect(err2).To(HaveOccurred())
			Expect(err2).To(MatchError(ContainSubstring("UNIQUE constraint failed: user.employee_id")))
		})

		It("throws an error when it cardID already exists", func() {
			user1 := getUserOne()
			user2 := getUserOne()
			err1 := userRepository.CreateUser(user1)
			err2 := userRepository.CreateUser(user2)

			Expect(err1).To(BeNil())
			Expect(err2).To(HaveOccurred())
			Expect(err2).To(MatchError(ContainSubstring("UNIQUE constraint failed: user.card_id")))
		})

		//It("throws an error when it cardID is not a alphanumeric string", func() {
		//	user := Models.User{
		//		EmployeeID: 2,
		//		CardID:     "bbbb",
		//		Name:       "Max Power",
		//		Email:      "max.power@gmail.com",
		//		Phone:      "09716244907",
		//		Pin:        "1234",
		//		ConfirmPin: "1234",
		//	}
		//
		//	err := userRepository.CreateUser(user)
		//
		//	Expect(err).To(HaveOccurred())
		//	Expect(err).To(MatchError(ContainSubstring("UNIQUE constraint failed: user.card_id")))
		//})
	})

	Context("Get User By ID", func() {
		It("gets a user by ID", func() {

		})
	})
})

func GetDatabase() {
	var err error
	db, err = gorm.Open("sqlite3", "UserRepositoryTest.db")
	if err != nil {
		fmt.Println("Failed to connect to database: ", err)
	}
}

func getUserOne() Models.User {
	user := Models.User{
		EmployeeID: 2,
		CardID:     "r7jTG7dqBy5wGO4L",
		Name:       "Max Power",
		Email:      "max.power@gmail.com",
		Phone:      "09716244907",
		Pin:        "1234",
		ConfirmPin: "1234",
	}
	return user
}

func getUserTwo() Models.User {
	user := Models.User{
		EmployeeID: 4,
		CardID:     "r7jTG7dqBy5wRK70",
		Name:       "Maxeen Power",
		Email:      "maxeen.power@gmail.com",
		Phone:      "09716244907",
		Pin:        "5432",
		ConfirmPin: "5432",
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
