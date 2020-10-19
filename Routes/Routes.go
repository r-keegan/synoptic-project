package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/r-keegan/synoptic-project/Controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// user routes
	r.GET("user", Controllers.GetUsers)
	r.POST("user", Controllers.CreateUser)
	r.GET("user/:id", Controllers.GetUserByID)

	return r
}
