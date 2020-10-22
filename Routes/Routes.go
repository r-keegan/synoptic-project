package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/r-keegan/synoptic-project/Config"
	"github.com/r-keegan/synoptic-project/Controllers"
	"github.com/r-keegan/synoptic-project/Repository"
	"github.com/r-keegan/synoptic-project/Services"
	"github.com/r-keegan/synoptic-project/Services/Session"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	userRepository := Repository.UserRepository{DB: Config.DB}
	userService := Services.UserService{UserRepository: userRepository}
	sessionService := Session.SessionService{MaxSessionLengthInSeconds: Config.MaxSessionLengthInSeconds}
	membershipController := Controllers.MembershipController{UserService: userService, SessionService: sessionService}

	r.GET("cardPresented/:id", membershipController.CardPresented) // tested
	r.GET("user/auth", membershipController.UserAuthenticate)      // tested
	r.GET("logout/:id", membershipController.LogOut)               // tested
	r.POST("user", membershipController.CreateUser)                // tested
	r.GET("balance", membershipController.GetBalance)              // tested
	r.PUT("purchase", membershipController.Purchase)               // tested
	r.PUT("topup", membershipController.TopUp)                     // tested

	return r
}
