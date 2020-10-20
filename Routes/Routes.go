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
	userController := Controllers.UserController{UserService: userService}
	// user routes
	//r.GET("user", Controllers.GetUsers)
	r.POST("user", userController.CreateUser)
	//r.GET("user/:id", Controllers.GetUserByID)

	return r
}
