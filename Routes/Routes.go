package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/r-keegan/synoptic-project/Config"
	"github.com/r-keegan/synoptic-project/Controllers"
	"github.com/r-keegan/synoptic-project/Services"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	userService := Services.UserService{DB: Config.DB}
	userController := Controllers.UserController{UserService: userService}
	// user routes
	//r.GET("user", Controllers.GetUsers)
	r.POST("user", userController.CreateUser)
	//r.GET("user/:id", Controllers.GetUserByID)

	// card routes
	//r.GET("card/:cardID", Controllers.GetCard)
	//r.POST("card", Controllers.CreateCard)
	//r.PUT("card/:cardID", Controllers.UpdateCard)
	//r.PUT("card/:cardID", Controllers.UpdateBalance)
	//r.DELETE("card/:cardID", Controllers.DeleteCard)

	return r
}
