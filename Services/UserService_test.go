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
	RunSpecs(t, "User Services Suite")
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
		_ = userRepository.CreateUser(existingUserWithAPositiveBalance())
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

		It("throws error when there attempts create the same user twice", func() {
			user := getUserOne()

			err := userService.CreateUser(user)
			Expect(err).ShouldNot(HaveOccurred())

			err = userService.CreateUser(user)
			Expect(err).Should(HaveOccurred())
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
			invalidUser := Models.CreateUser{
				EmployeeID: 1,
				Name:       "Max Power",
				Email:      "",
				Phone:      "09716244907",
				Pin:        "1234",
			}

			err := userService.CreateUser(invalidUser)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Required email")))
		})

		It("throws error when email is invalid", func() {
			invalidUser := Models.CreateUser{
				EmployeeID: 1,
				Name:       "Max Power",
				Email:      "max.powergmail.com",
				Phone:      "09716244907",
				Pin:        "1234",
			}

			err := userService.CreateUser(invalidUser)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Invalid email")))
		})

		It("throws error when phone is missing", func() {
			invalidUser := Models.CreateUser{
				EmployeeID: 1,
				Name:       "Max Power",
				Email:      "max.power@gmail.com",
				Phone:      "",
				Pin:        "1234",
			}

			err := userService.CreateUser(invalidUser)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Required phone")))
		})

		It("throws error when pin does not contain four numbers", func() {
			invalidUser := Models.CreateUser{
				EmployeeID: 1,
				Name:       "Max Power",
				Email:      "max.power@gmail.com",
				Phone:      "09716244907",
				Pin:        "06",
			}

			err := userService.CreateUser(invalidUser)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Invalid pin")))
		})

		It("throws error when pin does not contain four numbers", func() {
			invalidUser := Models.CreateUser{
				EmployeeID: 1,
				Name:       "Max Power",
				Email:      "max.power@gmail.com",
				Phone:      "09716244907",
				Pin:        "06ab",
			}

			err := userService.CreateUser(invalidUser)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Invalid pin")))
		})

		It("throws error cardID is not 16 characters long", func() {
			invalidUser := Models.CreateUser{
				EmployeeID: 1,
				Name:       "Max Power",
				Email:      "max.power@gmail.com",
				Phone:      "09716244907",
				Pin:        "2344",
				CardID:     "123412341234123",
			}

			err := userService.CreateUser(invalidUser)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Invalid cardID")))
		})

		It("throws error cardID is not alphanumeric long", func() {
			invalidUser := Models.CreateUser{
				EmployeeID: 1,
				Name:       "Max Power",
				Email:      "max.power@gmail.com",
				Phone:      "09716244907",
				Pin:        "2344",
				CardID:     "123412341234123,",
			}

			err := userService.CreateUser(invalidUser)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Invalid cardID")))
		})
	})

	Context("Update User", func() {
		It("successfully updates a user", func() {
			userUpdate := existingUserWithAPositiveBalance()
			userUpdate.Phone = userUpdate.Phone + "1"

			err := userService.UpdateUser(userUpdate)
			Expect(err).ShouldNot(HaveOccurred())

			actualResult, _ := userRepository.GetUserByCardID(userUpdate.CardID)
			Expect(actualResult.Phone).To(Equal(userUpdate.Phone))
		})

		It("throws error when employeeID is missing", func() {
			invalidUser := Models.User{
				Name: "Max Power",
			}

			err := userService.UpdateUser(invalidUser)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Required employeeID")))
		})

		It("throws error when name us missing", func() {
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

		It("throws error when balance tries to fall below zero", func() {
			invalidUser := Models.User{
				EmployeeID: 1,
				Name:       "Max Power",
				Email:      "max.power@gmail.com",
				Phone:      "09716244907",
				Pin:        "0634",
				Balance:    -30,
			}

			err := userService.UpdateUser(invalidUser)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Insufficient funds")))
		})
	})

	Context("Get employee by cardID", func() {
		It("gets a user by cardID", func() {
			user := getUserOne()
			err := userService.CreateUser(user)
			Expect(err).ShouldNot(HaveOccurred())

			actualUser, err := userService.GetEmployeeByCardID(user.CardID)

			Expect(err).ShouldNot(HaveOccurred())
			compareCreateUserToUser(user, actualUser)
		})
	})

	Context("Authenticate", func() {
		It("returns true if user is successfully authenticated", func() {
			user := getUserOne()
			err := userService.CreateUser(user)
			Expect(err).ShouldNot(HaveOccurred())

			userToAuth := Models.AuthenticatedRequest{
				CardID: "17jTG7dqBy5wGO4L",
				Pin:    "1234",
			}

			actualResult := userService.Authenticate(userToAuth)
			Expect(actualResult).To(BeTrue())
		})

		It("returns false if user fails to be authenticated", func() {
			user := getUserOne()
			err := userService.CreateUser(user)
			Expect(err).ShouldNot(HaveOccurred())

			userToAuth := Models.AuthenticatedRequest{
				CardID: "17jTG7dqBy5wGO4L",
				Pin:    "1",
			}

			actualResult := userService.Authenticate(userToAuth)
			Expect(actualResult).To(BeFalse())
		})
	})

	Context("Balance", func() {
		It("returns the balance", func() {
			user := existingUserWithAPositiveBalance()

			actualResult, err := userService.Balance(user.CardID, user.Pin)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(user.Balance).To(Equal(actualResult))
		})

		It("returns error if user pin incorrect", func() {
			existingUserWithAPositiveBalance()

			userToAuth := Models.AuthenticatedRequest{
				CardID: "moneymoneymoney1",
				Pin:    "1",
			}

			_, err := userService.Balance(userToAuth.CardID, userToAuth.Pin)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("user and pin mismatch")))
		})

		It("returns error if balance is unable to be obtained", func() {
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

			_, err2 := userService.Balance(user2.CardID, user2.Pin)
			Expect(err2).To(HaveOccurred())
		})
	})

	Context("Purchase", func() {
		It("successfully purchases an item and returns new balance", func() {
			user := existingUserWithAPositiveBalance()

			balance, err := userService.Purchase(user.CardID, user.Pin, 50)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(balance).To(Equal(999999950))
		})

		It("returns error if user pin incorrect", func() {
			existingUserWithAPositiveBalance()

			purchaseRequest := Models.PurchaseRequest{
				CardID: "moneymoneymoney1",
				Pin:    "1",
			}

			_, err := userService.Purchase(purchaseRequest.CardID, purchaseRequest.Pin, 50)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Unable to make purchase")))
		})

		It("returns error if purchase price is greater than balance", func() {
			user := existingUserWithAPositiveBalance()

			_, err := userService.Purchase(user.CardID, user.Pin, 100000000000)

			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Unable to make purchase")))
		})

		It("returns error if purchase price is a negative number", func() {
			user := existingUserWithAPositiveBalance()

			_, err := userService.Purchase(user.CardID, user.Pin, -30)

			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Purchase Amount is not valid")))
		})
	})

	Context("Topup", func() {
		It("successfully allows top up and returns new balance", func() {
			user := existingUserWithAPositiveBalance()

			balance, err := userService.TopUp(user.CardID, user.Pin, 50)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(balance).To(Equal(1000000050))
		})

		It("returns error if user pin incorrect", func() {
			existingUserWithAPositiveBalance()

			purchaseRequest := Models.TopUpRequest{
				CardID: "moneymoneymoney1",
				Pin:    "1",
			}

			_, err := userService.TopUp(purchaseRequest.CardID, purchaseRequest.Pin, 50)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("Unable to topup")))
		})

		It("returns error if top up amount is a negative number", func() {
			user := existingUserWithAPositiveBalance()

			_, err := userService.TopUp(user.CardID, user.Pin, -30)

			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("TopUp Amount is not valid")))
		})

		It("returns error if top up amount is zero", func() {
			user := existingUserWithAPositiveBalance()

			_, err := userService.TopUp(user.CardID, user.Pin, 0)

			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(ContainSubstring("TopUp Amount is not valid")))
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
		CardID:     "17jTG7dqBy5wGO4L",
		Email:      "max.power@gmail.com",
		Phone:      "09716244907",
		Pin:        "1234",
	}
	return user
}

func existingUserWithAPositiveBalance() Models.User {
	user := Models.User{
		ID:         1,
		EmployeeID: 1,
		Name:       "Richie Rich",
		CardID:     "moneymoneymoney1",
		Email:      "richierich@example.com",
		Phone:      "09716244907",
		Pin:        "4321",
		Balance:    1000000000,
	}
	return user
}

func compareCreateUserToUser(expectedUser Models.CreateUser, actualUser Models.User) {
	Expect(actualUser.EmployeeID).To(Equal(expectedUser.EmployeeID))
	Expect(actualUser.Name).To(Equal(expectedUser.Name))
	Expect(actualUser.Phone).To(Equal(expectedUser.Phone))
	Expect(actualUser.Pin).To(Equal(expectedUser.Pin))
}
