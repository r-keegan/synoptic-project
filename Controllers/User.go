package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/r-keegan/synoptic-project/Models"
	"github.com/r-keegan/synoptic-project/Services"
	"net/http"
)

func GetUsers(c *gin.Context) {
	var user []Models.User

	err := Services.GetAllUsers(&user)
	if err != nil {
		fmt.Println("Could not get all users: %")
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func CreateUser(c *gin.Context) {
	var user Models.User

	err := CreateUserByUserModel(user)
	if err != nil {
		fmt.Println("Could not create user: ", err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func CreateUserByUserModel(user Models.User) error {
	err := Services.CreateUser(&user)
	return err
}

func GetUserByID(c *gin.Context) {
	var user Models.User

	// gin framework finds the first JSON
	// parameter labelled "id"
	id := c.Params.ByName("id")
	err := Services.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
