package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/r-keegan/synoptic-project/Config"
	"github.com/r-keegan/synoptic-project/Models"
	"github.com/r-keegan/synoptic-project/Repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var (
	db             *gorm.DB
	router         *gin.Engine
	w              *httptest.ResponseRecorder
	userRepository Repository.UserRepository
)

const SessionTimeoutInSeconds = 1

func TestServices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Services Suite")
}

var _ = Describe("Integration test", func() {

	BeforeSuite(func() {
		GetTestDatabase()
		Config.MaxSessionLengthInSeconds = SessionTimeoutInSeconds
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

			Expect(w.Body.String()).To(Equal("Welcome Maxeen Power"))
		})

		It("responds with an error, need to register card", func() {

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

		It("user attempts to log out when their session has timed out", func() {
			requestBody := fmt.Sprintf(`{"cardID":"%s","pin":"%s"}`, existingUser().CardID, existingUser().Pin)
			req, err := http.NewRequest("GET", "/user/auth", strings.NewReader(requestBody))
			router.ServeHTTP(w, req)
			time.Sleep((SessionTimeoutInSeconds + 1) * time.Second) //Wait for session to timeout

			req, err = http.NewRequest("GET", fmt.Sprintf("/logout/%s", existingUser().CardID), nil)
			w2 := httptest.NewRecorder()
			router.ServeHTTP(w2, req)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(w2.Code).To(Equal(200))
			Expect(w2.Body.String()).To(Equal("User does not have a session"))
		})

		It("user attempts to log out when they do not have a session", func() {
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
			req, err := http.NewRequest("PUT", "/topup", strings.NewReader(requestBody))
			router.ServeHTTP(w, req)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Your balance is: 200"))
		})

		It("responds with an error, when unable to topup", func() {
			nonExistantCardID := existingUser().CardID
			requestBody := fmt.Sprintf(`{"cardID":"%s","pin":"%s", "amount":1}`, nonExistantCardID, existingUser().Pin+"2")
			req, err := http.NewRequest("PUT", "/topup", strings.NewReader(requestBody))
			router.ServeHTTP(w, req)

			Expect(err).ToNot(HaveOccurred())
			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Unable to topup"))
		})
	})

	Context("Purchase", func() {
		It("responds with 200 and informs user of their balance", func() {
			requestBody := fmt.Sprintf(`{"cardID":"%s","pin":"%s", "amount":100}`, existingUser().CardID, existingUser().Pin)
			req, err := http.NewRequest("PUT", "/purchase", strings.NewReader(requestBody))
			router.ServeHTTP(w, req)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Your balance is: 0"))
		})

		It("responds with an error, when purchase is greater than their existing balance", func() {
			requestBody := fmt.Sprintf(`{"cardID":"%s","pin":"%s", "amount":100000}`, existingUser().CardID, existingUser().Pin)
			req, err := http.NewRequest("PUT", "/purchase", strings.NewReader(requestBody))
			router.ServeHTTP(w, req)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("Unable to make purchase"))
		})
	})
})

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
