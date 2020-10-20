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

	r.GET("cardPresented/:id", membershipController.CardPresented)
	r.GET("user/auth", membershipController.UserAuthenticate)
	r.GET("Logout/:id", membershipController.LogOut)
	r.POST("user", membershipController.CreateUser)
	//r.GET("user/:id", Controllers.GetUserByID)

	return r
}
