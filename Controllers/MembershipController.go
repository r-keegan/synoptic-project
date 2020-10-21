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
	var authRequest Models.AuthenticatedRequest
	context.BindJSON(&authRequest)
	authenticationResult := c.UserService.Authenticate(authRequest)

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

func (c MembershipController) GetBalance(context *gin.Context) {
	var authRequest Models.AuthenticatedRequest
	context.BindJSON(&authRequest)
	balance, err := c.UserService.GetBalance(authRequest.CardID, authRequest.Pin)
	if err == nil {
		context.String(http.StatusOK, fmt.Sprintf("Your balance is: %v", balance))
	} else {
		context.String(http.StatusOK, fmt.Sprintf("Unable to provide balance"))
	}
}

func (c MembershipController) Purchase(context *gin.Context) {
	var purchaseRequest Models.PurchaseRequest
	context.BindJSON(&purchaseRequest)
	balance, err := c.UserService.Purchase(purchaseRequest.CardID, purchaseRequest.Pin, purchaseRequest.Amount)
	if err == nil {
		context.String(http.StatusOK, fmt.Sprintf("Your balance is: %v", balance))
	} else {
		context.String(http.StatusOK, fmt.Sprintf("Unable to make purchase: your balance is %v", balance))
	}
}

func (c MembershipController) TopUp(context *gin.Context) {
	var topUpRequest Models.TopUpRequest
	context.BindJSON(&topUpRequest)
	balance, err := c.UserService.TopUp(topUpRequest.CardID, topUpRequest.Pin, topUpRequest.Amount)
	if err == nil {
		context.String(http.StatusOK, fmt.Sprintf("Your balance is: %v", balance))
	} else {
		context.String(http.StatusOK, fmt.Sprintf("Unable to topup: your balance is %v", balance))
	}
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
