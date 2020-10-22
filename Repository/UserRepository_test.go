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

		It("throws an error when an employeeID already exists", func() {
			user1 := getUserOne()
			user2 := Models.User{
				EmployeeID: 2,
				CardID:     "b51TG7dqBy5wGO4L",
				Name:       "Max Power",
				Email:      "max.power@gmail.com",
				Phone:      "09716244907",
				Pin:        "1234",
			}
			err1 := userRepository.CreateUser(user1)
			err2 := userRepository.CreateUser(user2)

			Expect(err1).To(BeNil())
			Expect(err2).To(HaveOccurred())
			Expect(err2).To(MatchError(ContainSubstring("UNIQUE constraint failed: user.employee_id")))
		})

		It("throws an error when cardID already exists", func() {
			user1 := getUserOne()
			user2 := getUserOne()
			err1 := userRepository.CreateUser(user1)
			err2 := userRepository.CreateUser(user2)

			Expect(err1).To(BeNil())
			Expect(err2).To(HaveOccurred())
			Expect(err2).To(MatchError(ContainSubstring("UNIQUE constraint failed: user.card_id")))
		})
	})

	Context("Get User", func() {
		It("successfully gets a user by their unique employeeID", func() {
			//Given I have a User
			user := getUserOne()

			// and I create a user
			err := userRepository.CreateUser(user)
			var foundUser Models.User

			// I can find the user by their unique employee ID
			foundUser, err = userRepository.GetUserByEmployeeID(2)
			Expect(err).ShouldNot(HaveOccurred())

			compare(foundUser, user)
		})

		It("throws an error when it is unable to find their unique employee ID", func() {
			// Given I have a User
			user := Models.User{
				EmployeeID: 78890,
				CardID:     "r7jTG7dqBy5wGO4L",
				Name:       "Max Power",
				Email:      "max.power@gmail.com",
				Phone:      "09716244907",
				Pin:        "1234",
			}
			// and I create a user
			err := userRepository.CreateUser(user)

			// I cannot find the user by their unique employee ID
			_, err = userRepository.GetUserByEmployeeID(2)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("record not found")))
		})

		It("successfully gets a user by their unique cardID", func() {
			//Given I have a User
			user := getUserOne()

			// and I create a user
			err := userRepository.CreateUser(user)
			var foundUser Models.User

			// I can find the user by their unique employee ID
			foundUser, err = userRepository.GetUserByCardID("r7jTG7dqBy5wGO4L")
			Expect(err).ShouldNot(HaveOccurred())

			compare(foundUser, user)
		})

		It("throws an error when it is unable to find their unique card ID", func() {
			// Given I have a User
			user := Models.User{
				EmployeeID: 2,
				CardID:     "r7jTG7dqBy5wGO4L",
				Name:       "Max Power",
				Email:      "max.power@gmail.com",
				Phone:      "09716244907",
				Pin:        "1234",
			}

			// and I create a user
			err := userRepository.CreateUser(user)

			// I cannot find the user by their unique employee ID
			_, err = userRepository.GetUserByCardID("")
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("record not found")))
		})
	})

	Context("Update user", func() {
		It("successfully updates user information when pin changes", func() {
			user := getUserOne()
			err := userRepository.CreateUser(user)

			user = Models.User{
				EmployeeID: 2,
				CardID:     "r7jTG7dqBy5wGO4L",
				Name:       "Max Power",
				Email:      "max.power@gmail.com",
				Phone:      "09716244907",
				Pin:        "7890",
			}
			err = userRepository.UpdateUser(user)
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("successfully updates user information when cardID changes", func() {
			user := getUserOne()
			err := userRepository.CreateUser(user)

			user = Models.User{
				EmployeeID: 2,
				CardID:     "ppp1G7dqBy5wGO4L",
				Name:       "Max Power",
				Email:      "max.power@gmail.com",
				Phone:      "09716244907",
				Pin:        "7890",
			}
			err = userRepository.UpdateUser(user)
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Delete user", func() {
		It("successfully deletes a user by their unique employeeID", func() {
			user := getUserOne()
			err := userRepository.CreateUser(user)
			err = userRepository.DeleteUserByEmployeeID(2)
			Expect(err).ShouldNot(HaveOccurred())

			_, err = userRepository.GetUserByEmployeeID(2)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("record not found")))
		})

		It("successfully deletes a user by their unique cardID", func() {
			user := getUserTwo()
			err := userRepository.CreateUser(user)
			err = userRepository.DeleteUserByCardID("r7jTG7dqBy5wRK70")
			Expect(err).ShouldNot(HaveOccurred())

			_, err = userRepository.GetUserByEmployeeID(2)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("record not found")))
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
	}
	return user
}

//Test Helpers
func compare(expectedUser Models.User, actualUser Models.User) {
	Expect(actualUser.EmployeeID).To(Equal(expectedUser.EmployeeID))
	Expect(actualUser.Name).To(Equal(expectedUser.Name))
	Expect(actualUser.Phone).To(Equal(expectedUser.Phone))
	Expect(actualUser.Pin).To(Equal(expectedUser.Pin))
}
