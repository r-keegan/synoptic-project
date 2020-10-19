package Controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/r-keegan/synoptic-project/Models"
	"github.com/r-keegan/synoptic-project/Services"
	"net/http"
)

//func GetUsers(c *gin.Context) {
//	var user []Models.User
//
//	err := Services.GetAllUsers(&user)
//	if err != nil {
//		fmt.Println("Could not get all users: %")
//		c.AbortWithStatus(http.StatusNotFound)
//	} else {
//		c.JSON(http.StatusOK, user)
//	}
//}
type UserController struct {
	UserService Services.UserService
}

func (s UserController) CreateUser(c *gin.Context) {
	user := mapGinContextToUser(c)

	err := s.UserService.CreateUser(user)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func mapGinContextToUser(c *gin.Context) Models.User {
	var user Models.User
	c.BindJSON(&user)
	return user
}

//
//func GetUserByID(c *gin.Context) {
//	var user Models.User
//
//	// gin framework finds the first JSON parameter labelled "id"
//	id := c.Params.ByName("id")
//	err := Services.GetUserByID(&user, id)
//	if err != nil {
//		c.AbortWithStatus(http.StatusNotFound)
//	} else {
//		c.JSON(http.StatusOK, user)
//	}
//}
