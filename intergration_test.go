package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/r-keegan/synoptic-project/Controllers"
	"github.com/r-keegan/synoptic-project/Models"
	"gopkg.in/go-playground/assert.v1"
	"net/http"
	"net/http/httptest"
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
		EmployeeID: 3,
		Name:       "Max Power",
		Email:      "max.power@gmail.com",
		Phone:      "09716244907",
		Pin:        1234,
	}

	Controllers.CreateUserByUserModel(user)
	req, _ := http.NewRequest("GET", "/user", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[{\"id\":1,\"employeeID\":3,\"name\":\"Max Power\",\"email\":\"max.power@gmail.com\",\"phone\":\"09716244907\",\"pin\":1234}]", w.Body.String())
}

func TestGetUsers_WhenMultipleUses(t *testing.T) {
	user := Models.User{
		ID:         1,
		EmployeeID: 3,
		Name:       "Max Power",
		Email:      "max.power@gmail.com",
		Phone:      "09716244907",
		Pin:        1234,
	}

	Controllers.CreateUserByUserModel(user)

	// change details to create a second user
	user.ID = 2
	user.EmployeeID = 7
	user.Pin = 5432
	Controllers.CreateUserByUserModel(user)

	req, _ := http.NewRequest("GET", "/user", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[{\"id\":1,\"employeeID\":3,\"name\":\"Max Power\",\"email\":\"max.power@gmail.com\",\"phone\":\"09716244907\",\"pin\":1234},{\"id\":2,\"employeeID\":7,\"name\":\"Max Power\",\"email\":\"max.power@gmail.com\",\"phone\":\"09716244907\",\"pin\":5432}]", w.Body.String())
}
