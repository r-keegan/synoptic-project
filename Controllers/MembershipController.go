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

var sessions = make(map[string]bool)

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
	_ = context.BindJSON(&createUser)
	err := c.UserService.CreateUser(createUser)
	if err != nil {
		context.String(http.StatusOK, fmt.Sprintf("Unable to create user: %v", err))
	} else {
		context.String(http.StatusOK, "User created")
	}
}

func (c MembershipController) UserAuthenticate(context *gin.Context) {
	var authRequest Models.AuthenticatedRequest
	_ = context.BindJSON(&authRequest)
	authenticationResult := c.UserService.Authenticate(authRequest)

	if authenticationResult {
		sessions[authRequest.CardID] = true
		context.String(http.StatusOK, "Log in successful")
	} else {
		context.String(http.StatusOK, "Log in failed")
	}
}

func (c MembershipController) LogOut(context *gin.Context) {
	cardId := context.Params.ByName("id")
	// TODO Invalidate session for card
	//todo create a test that tries to logout using a CardID that has never logged in
	if sessions[cardId] {
		sessions[cardId] = false
		context.String(http.StatusOK, "Goodbye")
	} else {
		context.String(http.StatusOK, "User does not have a session")
	}
}

func (c MembershipController) GetBalance(context *gin.Context) {
	var authRequest Models.AuthenticatedRequest
	_ = context.BindJSON(&authRequest)
	balance, err := c.UserService.Balance(authRequest.CardID, authRequest.Pin)
	if err == nil {
		context.String(http.StatusOK, fmt.Sprintf("Your balance is: %v", balance))
	} else {
		context.String(http.StatusOK, fmt.Sprintf("Unable to provide balance"))
	}
}

func (c MembershipController) Purchase(context *gin.Context) {
	var purchaseRequest Models.PurchaseRequest
	_ = context.BindJSON(&purchaseRequest)
	balance, err := c.UserService.Purchase(purchaseRequest.CardID, purchaseRequest.Pin, purchaseRequest.Amount)
	if err == nil {
		context.String(http.StatusOK, fmt.Sprintf("Your balance is: %v", balance))
	} else {
		context.String(http.StatusOK, fmt.Sprintf("Unable to make purchase"))
	}
}

func (c MembershipController) TopUp(context *gin.Context) {
	var topUpRequest Models.TopUpRequest
	_ = context.BindJSON(&topUpRequest)
	balance, err := c.UserService.TopUp(topUpRequest.CardID, topUpRequest.Pin, topUpRequest.Amount)
	if err == nil {
		context.String(http.StatusOK, fmt.Sprintf("Your balance is: %v", balance))
	} else {
		context.String(http.StatusOK, fmt.Sprintf("Unable to topup"))
	}
}
