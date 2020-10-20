package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/r-keegan/synoptic-project/Models"
	"github.com/r-keegan/synoptic-project/Services"
	"net/http"
)

type MembershipController struct {
	UserService Services.UserService
}

func (c MembershipController) CardPresented(context *gin.Context) {
	cardID := context.Params.ByName("id")

	user, err := c.UserService.GetEmployeeByCardID(cardID)
	if err != nil {
		context.String(http.StatusOK, "Card needs to be registered")
	} else {
		context.String(http.StatusOK, fmt.Sprintf("Welcome %s", user.Name))
	}
}

//func (c MembershipController) CreateUser(context *gin.Context) {
//	user := mapGinContextToUser(context)
//	err := c.UserService.CreateUser(user)
//	if err != nil {
//		context.AbortWithStatus(http.StatusInternalServerError)
//	} else {
//		context.JSON(http.StatusOK, user)
//	}
//}
//func mapGinContextToUser(c *gin.Context) Models.User {
//	var user Models.User
//	c.BindJSON(&user)
//	return user
//}

func (c MembershipController) CreateUser(context *gin.Context) {
	var createUser Models.CreateUser
	context.BindJSON(&createUser)
	err := c.UserService.CreateUser(createUser)
	if err != nil {
		context.String(http.StatusOK, fmt.Sprintf("Unable to create user: %v", err))
	} else {
		context.String(http.StatusOK, "User created")
	}
}

// TODO store logged-in somewhere eg: session
func (c MembershipController) UserAuthenticate(context *gin.Context) {
	var userAuth Models.UserAuth
	context.BindJSON(&userAuth)
	authenticationResult := c.UserService.Authenticate(userAuth)

	if authenticationResult {
		context.String(http.StatusOK, "Log in successful")
	} else {
		context.String(http.StatusOK, "Log in failed")
	}
}

func (c MembershipController) LogOut(context *gin.Context) {
	context.Params.ByName("id")
	// TODO Invalidate session for card
	// TODO If session is valid then say goodbye
	context.String(http.StatusOK, "Goodbye")
}

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
