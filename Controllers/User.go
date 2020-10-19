package Controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/r-keegan/synoptic-project/Models"
	"net/http"
)

func GetUsers(c *gin.Context) {
	var user []Models.User

	err := Models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func CreateUser(c *gin.Context) {

}

func GetUserByID(c *gin.Context) {

}
