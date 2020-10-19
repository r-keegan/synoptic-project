package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/r-keegan/synoptic-project/Controllers"
	"github.com/r-keegan/synoptic-project/Models"
	"gopkg.in/go-playground/assert.v1"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	db     *gorm.DB
	router *gin.Engine
	w      *httptest.ResponseRecorder
)

func TestMain(m *testing.M) {
	// setup server { database + router }
	db = GetDatabase()
	db.DropTableIfExists(&Models.User{})
	router = SetupRouterWithSuppliedDB(db)
	w = httptest.NewRecorder()

	// runs tests
	m.Run()
}

func TestGetUser_WhenNoUsers(t *testing.T) {
	req, _ := http.NewRequest("GET", "/user", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[]", w.Body.String())
}

func TestGetUsers_WhenOneUser(t *testing.T) {
	user := Models.User{
		ID:         1,
		EmployeeID: 2,
		Name:       "Max Power",
		Email:      "max.power@gmail.com",
		Phone:      "09716244907",
		Pin:        1234,
	}

	Controllers.CreateUserByUserModel(user)
	req, _ := http.NewRequest("GET", "/user", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[][{\"id\":1,\"employeeID\":2,\"name\":\"Max Power\",\"email\":\"max.power@gmail.com\",\"phone\":\"09716244907\",\"pin\":1234}]", w.Body.String())
}

func TestGetUsers_WhenMultipleUses(t *testing.T) {
	user := Models.User{
		ID:         3,
		EmployeeID: 4,
		Name:       "Max Power",
		Email:      "max.power@gmail.com",
		Phone:      "09716244907",
		Pin:        1234,
	}

	Controllers.CreateUserByUserModel(user)

	// change details to create a second user
	user.ID = 5
	user.EmployeeID = 6
	user.Pin = 5432
	Controllers.CreateUserByUserModel(user)

	req, _ := http.NewRequest("GET", "/user", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[][{\"id\":1,\"employeeID\":2,\"name\":\"Max Power\",\"email\":\"max.power@gmail.com\",\"phone\":\"09716244907\",\"pin\":1234}][{\"id\":3,\"employeeID\":4,\"name\":\"Max Power\",\"email\":\"max.power@gmail.com\",\"phone\":\"09716244907\",\"pin\":1234},{\"id\":5,\"employeeID\":6,\"name\":\"Max Power\",\"email\":\"max.power@gmail.com\",\"phone\":\"09716244907\",\"pin\":5432}]", w.Body.String())
}

func TestCreateUser(t *testing.T) {
	requestBody := fmt.Sprintf(
		`{"employeeId":5,
				"name":"Maximum Power",
				"email":"maximum.power@gmail.com",
				"phone":"0777000000",
				"pin":1234
				}`)
	req, _ := http.NewRequest("POST", "/user", strings.NewReader(requestBody))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[][{\"id\":9,\"employeeID\":5,\"name\":\"Maximum Power\",\"email\":\"maximum.power@gmail.com\",\"phone\":\"0777000000\",\"pin\":1234}]", w.Body.String())
}

func TestGetUserByID(t *testing.T) {
	//user := Models.User{
	//	ID:         7,
	//	EmployeeID: 8,
	//	Name:       "Maxeen Power",
	//	Email:      "maxeen.power@gmail.com",
	//	Phone:      "0900111222",
	//	Pin:        1234,
	//}
	//
	//Controllers.CreateUserByUserModel(user)
	req, _ := http.NewRequest("GET", "/user/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"id\":1,\"employeeID\":2,\"name\":\"Max Power\",\"email\":\"max.power@gmail.com\",\"phone\":\"09716244907\",\"pin\":1234}", w.Body.String())
}
