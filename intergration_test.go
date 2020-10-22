package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/r-keegan/synoptic-project/Models"
	"github.com/r-keegan/synoptic-project/Repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	db             *gorm.DB
	router         *gin.Engine
	w              *httptest.ResponseRecorder
	userRepository Repository.UserRepository
)

func TestServices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Services Suite")
}

//goland:noinspection ALL
var _ = Describe("Integration test", func() {

	BeforeSuite(func() {
		GetTestDatabase()
		router = SetupRouterWithSuppliedDB(db)
	})

	BeforeEach(func() {
		// setup database
		db.DropTableIfExists(&Models.User{})
		db.AutoMigrate(&Models.User{})

		userRepository = Repository.UserRepository{DB: db}
		_ = userRepository.CreateUser(existingUser())
		w = httptest.NewRecorder()

	})

	AfterEach(func() {
		db.DropTableIfExists(&Models.User{})
	})

	Context("Create User", func() {
		It("responds with 200 and successfully creates a user", func() {
			requestBody := getUserOneRequestBody()
			req, err := http.NewRequest("POST", "/user", strings.NewReader(requestBody))
			router.ServeHTTP(w, req)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("User created"))
		})

		It("responds with an error, unable to create user", func() {
			requestBody := getUserOneRequestBody()
			req, _ := http.NewRequest("POST", "/user", strings.NewReader(requestBody))
			router.ServeHTTP(w, req)

			requestBody2 := getUserOneRequestBody()
			req, err := http.NewRequest("POST", "/user", strings.NewReader(requestBody2))
			w2 := httptest.NewRecorder()
			router.ServeHTTP(w2, req)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(w2.Code).To(Equal(200))
			Expect(w2.Body.String()).To(Equal("Unable to create user: UNIQUE constraint failed: user.card_id"))
		})
	})

	Context("CardPresented", func() {
		It("responds with 200 and displays welcome message", func() {
			req, err := http.NewRequest("GET", fmt.Sprintf("/cardPresented/%s", existingUser().CardID), nil)
			router.ServeHTTP(w, req)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(w.Code).To(Equal(200))
			//goland:noinspection SpellCheckingInspection
			Expect(w.Body.String()).To(Equal("Welcome Maxeen Power"))
		})

		It("responds with an error, unable to register card", func() {
			//goland:noinspection SpellCheckingInspection
			req, err := http.NewRequest("GET", "/cardPresented/nocard", nil)
			router.ServeHTTP(w, req)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Card needs to be registered"))
		})
	})

	Context("LogOut", func() {
		It("user successfully logs out", func() {
			requestBody := fmt.Sprintf(`{"cardID":"%s","pin":"%s"}`, existingUser().CardID, existingUser().Pin)
			req, err := http.NewRequest("GET", "/user/auth", strings.NewReader(requestBody))
			router.ServeHTTP(w, req)

			req, err = http.NewRequest("GET", fmt.Sprintf("/logout/%s", existingUser().CardID), nil)
			w2 := httptest.NewRecorder()
			router.ServeHTTP(w2, req)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(w2.Code).To(Equal(200))
			Expect(w2.Body.String()).To(Equal("Goodbye"))
		})

		It("user attempts to logs out when they do not have a session", func() {
			req, err := http.NewRequest("GET", fmt.Sprintf("/logout/%s", existingUser().CardID), nil)
			router.ServeHTTP(w, req)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("User does not have a session"))
		})
	})

	Context("UserAuthenticate", func() {
		It("responds with 200 and informs user, log in successful", func() {
			requestBody := fmt.Sprintf(`{"cardID":"%s","pin":"%s"}`, existingUser().CardID, existingUser().Pin)
			req, err := http.NewRequest("GET", "/user/auth", strings.NewReader(requestBody))
			router.ServeHTTP(w, req)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Log in successful"))
		})

		It("responds with an error, informs user log in was unsuccessful", func() {
			requestBody := `{"cardID":"mystery-card","pin":"mystery-pin"}`
			req, err := http.NewRequest("GET", "/user/auth", strings.NewReader(requestBody))
			router.ServeHTTP(w, req)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Log in failed"))
		})
	})

	Context("Balance", func() {
		It("responds with 200 and informs user of their balance", func() {
			requestBody := fmt.Sprintf(`{"cardID":"%s","pin":"%s"}`, existingUser().CardID, existingUser().Pin)
			req, err := http.NewRequest("GET", "/balance", strings.NewReader(requestBody))
			router.ServeHTTP(w, req)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal(fmt.Sprintf("Your balance is: %v", existingUser().Balance)))
		})

		It("responds with an error, user unable to see balance", func() {
			requestBody := `{"cardID":"mystery-card","pin":"mystery-pin"}`
			req, err := http.NewRequest("GET", "/balance", strings.NewReader(requestBody))
			router.ServeHTTP(w, req)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Unable to provide balance"))
		})
	})

	Context("Topup", func() {
		It("responds with 200 when user tops up and informs user of their balance", func() {
			// existing user already has a balance of 100
			requestBody := fmt.Sprintf(`{"cardID":"%s","pin":"%s", "amount":100}`, existingUser().CardID, existingUser().Pin)
			//goland:noinspection ALL
			req, err := http.NewRequest("GET", "/topup", strings.NewReader(requestBody))
			router.ServeHTTP(w, req)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Your balance is: 200"))
		})

		It("responds with an error, when unable to topup", func() {
			nonExistantCardID := existingUser().CardID
			requestBody := fmt.Sprintf(`{"cardID":"%s","pin":"%s", "amount":1}`, nonExistantCardID, existingUser().Pin+"2")
			req, err := http.NewRequest("GET", "/topup", strings.NewReader(requestBody))
			router.ServeHTTP(w, req)

			Expect(err).ToNot(HaveOccurred())
			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Unable to topup"))
		})
	})

	Context("Purchase", func() {
		It("responds with 200 and informs user of their balance", func() {
			requestBody := fmt.Sprintf(`{"cardID":"%s","pin":"%s", "amount":100}`, existingUser().CardID, existingUser().Pin)
			req, err := http.NewRequest("GET", "/purchase", strings.NewReader(requestBody))
			router.ServeHTTP(w, req)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Your balance is: 0"))
		})

		It("responds with an error, when purchase is greater than their existing balance", func() {
			requestBody := fmt.Sprintf(`{"cardID":"%s","pin":"%s", "amount":100000}`, existingUser().CardID, existingUser().Pin)
			req, err := http.NewRequest("GET", "/purchase", strings.NewReader(requestBody))
			router.ServeHTTP(w, req)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Unable to make purchase"))
		})
	})
})

//func TestGetUser_WhenNoUsers(t *testing.T) {
//	req, _ := http.NewRequest("GET", "/user", nil)
//	router.ServeHTTP(w, req)
//
//	assert.Equal(t, 200, w.Code)
//	assert.Equal(t, "[]", w.Body.String())
//}

//func TestGetUsers_WhenOneUser(t *testing.T) {
//	//card := Models.Card{
//	//	CardID:  "r7jTG7dqBy5wGO4L",
//	//	Balance: 100,
//	//}
//	user := Models.User{
//		ID:         1,
//		EmployeeID: 2,
//		Name:       "Max Power",
//		Email:      "max.power@gmail.com",
//		Phone:      "09716244907",
//		Pin:        1234,
//	}
//
//	Services.UserRepository(user)
//	req, _ := http.NewRequest("GET", "/user", nil)
//	router.ServeHTTP(w, req)
//
//	assert.Equal(t, 200, w.Code)
//	assert.Equal(t, "[{\"id\":1,\"employeeID\":2,\"name\":\"Max Power\",\"email\":\"max.power@gmail.com\",\"phone\":\"09716244907\",\"pin\":1234}]", w.Body.String())
//}

//
//func TestGetUsers_WhenMultipleUses(t *testing.T) {
//	user := Models.User{
//		ID:         3,
//		EmployeeID: 4,
//		Name:       "Max Power",
//		Email:      "max.power@gmail.com",
//		Phone:      "09716244907",
//		Pin:        1234,
//	}
//
//	Controllers.CreateUserByUserModel(user)
//
//	// change details to create a second user
//	user.ID = 5
//	user.EmployeeID = 6
//	user.Pin = 5432
//	Controllers.CreateUserByUserModel(user)
//
//	req, _ := http.NewRequest("GET", "/user", nil)
//	router.ServeHTTP(w, req)
//	assert.Equal(t, 200, w.Code)
//	assert.Equal(t, "[{\"id\":3,\"employeeID\":4,\"name\":\"Max Power\",\"email\":\"max.power@gmail.com\",\"phone\":\"09716244907\",\"pin\":1234},{\"id\":5,\"employeeID\":6,\"name\":\"Max Power\",\"email\":\"max.power@gmail.com\",\"phone\":\"09716244907\",\"pin\":5432}]", w.Body.String())
//}

//func TestCreateUser_whenThatUserAlreadyExists(t *testing.T) {

//}

//func TestGetUserByID(t *testing.T) {
//	// given I have a user
//	user := newUser()
//	repository.CreateUser(user)
//	//when `i get
//
//	req, _ := http.NewRequest("GET", "/user/r7jTG7dqBy5wGO4L", nil)
//	router.ServeHTTP(w, req)
//	assert.Equal(t, 200, w.Code)
//	assert.Equal(t, "", w.Body.String())
//}

func getUserOneRequestBody() (requestBody string) {
	return `{"employeeID":90,"cardID":"2111TG7dqBy5wGO4","name":"Paul R","email":"paul.harris1@gmail.com","phone":"0799007","pin":"1234","balance":0}`
}

func GetTestDatabase() {
	var err error
	db, err = gorm.Open("sqlite3", "IntegrationTest.db")
	if err != nil {
		fmt.Println("Failed to connect to database: ", err)
	}
}

func existingUser() Models.User {
	user := Models.User{
		EmployeeID: 4,
		CardID:     "B7jTG7dqBy5wRK70",
		Name:       "Maxeen Power",
		Email:      "maxeen.power@gmail.com",
		Phone:      "09716244907",
		Pin:        "5432",
		Balance:    100,
	}
	return user
}
