package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
