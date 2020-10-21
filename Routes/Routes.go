package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/r-keegan/synoptic-project/Config"
	"github.com/r-keegan/synoptic-project/Controllers"
	"github.com/r-keegan/synoptic-project/Repository"
	"github.com/r-keegan/synoptic-project/Services"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	userRepository := Repository.UserRepository{DB: Config.DB}
	userService := Services.UserService{UserRepository: userRepository}
	membershipController := Controllers.MembershipController{UserService: userService}

	r.GET("cardPresented/:id", membershipController.CardPresented) // tested
	r.GET("user/auth", membershipController.UserAuthenticate)      // tested
	r.GET("logout/:id", membershipController.LogOut)               // tested
	r.POST("user", membershipController.CreateUser)                // tested
	r.GET("balance", membershipController.GetBalance)              // tested
	r.GET("purchase", membershipController.Purchase)               // tested
	r.GET("topup", membershipController.TopUp)                     // tested

	return r
}
