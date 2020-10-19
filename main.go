package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/r-keegan/synoptic-project/Config"
	"github.com/r-keegan/synoptic-project/Models"
	"github.com/r-keegan/synoptic-project/Routes"
)

func main() {
	r := SetUpRouter()
	r.Run()
}

func SetupRouterWithSuppliedDB(db *gorm.DB) *gin.Engine {
	Config.DB = db
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	return r
}

func SetUpRouter() *gin.Engine {
	db := GetDatabase()
	return SetupRouterWithSuppliedDB(db)
}

func GetDatabase() (db *gorm.DB) {
	db, err := gorm.Open("sqlite3", "main.db")
	if err != nil {
		fmt.Println("Failed to connect to database: ", err)
	}
	return db
}
