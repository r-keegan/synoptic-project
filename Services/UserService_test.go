package Services_test

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/r-keegan/synoptic-project/Models"
	"github.com/r-keegan/synoptic-project/Repository"
	. "github.com/r-keegan/synoptic-project/Services"
	"testing"
)

var db *gorm.DB
var userRepository Repository.UserRepository
var userService UserService

func TestServices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Services Suite")
}

var _ = Describe("UserService", func() {
	BeforeEach(func() {
		GetDatabase()
		// setup database
		db.DropTableIfExists(&Models.User{})
		db.AutoMigrate(&Models.User{})

		// set up user service
		userRepository = Repository.UserRepository{DB: db}
		userService = UserService{UserRepository: userRepository}
	})

	AfterEach(func() {
		db.DropTableIfExists(&Models.User{})
	})

	Context("Create User", func() {
		It("successfully creates a user", func() {
			user := getUserOne()

			err := userService.CreateUser(user)
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("throws error when employeeID missing", func() {
			invalidUser := Models.CreateUser{
				Name: "Max Power",
			}

			err := userService.CreateUser(invalidUser)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Required employeeID")))
		})

		It("throws error when name missing", func() {
			invalidUser := Models.CreateUser{
				EmployeeID: 1,
				Name:       "",
			}

			err := userService.CreateUser(invalidUser)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Required name")))
		})

		It("throws error when email missing", func() {
			invalidUser := Models.User{
				EmployeeID: 1,
				Name:       "Max Power",
				Email:      "",
				Phone:      "09716244907",
				Pin:        "1234",
			}

			err := userService.Validate(invalidUser, "update")
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Required email")))
		})

		It("throws error when email is invalid", func() {
			invalidUser := Models.User{
				EmployeeID: 1,
				Name:       "Max Power",
				Email:      "max.powergmail.com",
				Phone:      "09716244907",
				Pin:        "1234",
			}

			err := userService.Validate(invalidUser, "update")
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Invalid email")))
		})

		It("throws error when phone is missing", func() {
			invalidUser := Models.User{
				EmployeeID: 1,
				Name:       "Max Power",
				Email:      "max.power@gmail.com",
				Phone:      "",
				Pin:        "1234",
			}

			err := userService.Validate(invalidUser, "update")
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Required phone")))
		})

		It("throws error when pin does not contain four numbers", func() {
			invalidUser := Models.User{
				EmployeeID: 1,
				Name:       "Max Power",
				Email:      "max.power@gmail.com",
				Phone:      "09716244907",
				Pin:        "06",
			}

			err := userService.Validate(invalidUser, "update")
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Invalid pin")))
		})
	})

	Context("Update User", func() {
		It("successfully updates a user", func() {
			user := getUserOne()
			err := userService.CreateUser(user)
			Expect(err).ShouldNot(HaveOccurred())

			user2 := Models.User{
				EmployeeID: 2,
				CardID:     "123",
				Name:       "Max Power",
				Email:      "max.power@gmail.com",
				Phone:      "09716244907",
				Pin:        "5432",
				Balance:    0,
			}
			err = userService.UpdateUser(user2)
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("throws error when employeeID missing", func() {
			invalidUser := Models.User{
				Name: "Max Power",
			}

			err := userService.UpdateUser(invalidUser)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Required employeeID")))
		})

		It("throws error when name missing", func() {
			invalidUser := Models.User{
				EmployeeID: 1,
				Name:       "",
			}

			err := userService.UpdateUser(invalidUser)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Required name")))
		})

		It("throws error when email missing", func() {
			invalidUser := Models.User{
				EmployeeID: 1,
				Name:       "Max Power",
				Email:      "",
				Phone:      "09716244907",
				Pin:        "1234",
			}

			err := userService.UpdateUser(invalidUser)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Required email")))
		})

		It("throws error when email is invalid", func() {
			invalidUser := Models.User{
				EmployeeID: 1,
				Name:       "Max Power",
				Email:      "max.powergmail.com",
				Phone:      "09716244907",
				Pin:        "1234",
			}

			err := userService.UpdateUser(invalidUser)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Invalid email")))
		})

		It("throws error when phone is missing", func() {
			invalidUser := Models.User{
				EmployeeID: 1,
				Name:       "Max Power",
				Email:      "max.power@gmail.com",
				Phone:      "",
				Pin:        "1234",
			}

			err := userService.UpdateUser(invalidUser)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Required phone")))
		})

		It("throws error when pin does not contain four numbers", func() {
			invalidUser := Models.User{
				EmployeeID: 1,
				Name:       "Max Power",
				Email:      "max.power@gmail.com",
				Phone:      "09716244907",
				Pin:        "06",
			}

			err := userService.UpdateUser(invalidUser)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Invalid pin")))
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

func getUserOne() Models.CreateUser {
	user := Models.CreateUser{
		EmployeeID: 2,
		Name:       "Max Power",
		CardID:     "r7jTG7dqBy5wGO4L",
		Email:      "max.power@gmail.com",
		Phone:      "09716244907",
		Pin:        "1234",
	}
	return user
}

func getUserTwo() Models.User {
	user := Models.User{
		EmployeeID: 4,
		Name:       "Maxeen Power",
		Email:      "maxeen.power@gmail.com",
		Phone:      "09716244907",
		Pin:        "1234",
		Balance:    0,
	}
	return user
}

func compare(expectedUser Models.User, actualUser Models.User) {
	Expect(actualUser.EmployeeID).To(Equal(expectedUser.EmployeeID))
	Expect(actualUser.Name).To(Equal(expectedUser.Name))
	Expect(actualUser.Phone).To(Equal(expectedUser.Phone))
	Expect(actualUser.Pin).To(Equal(expectedUser.Pin))
}
