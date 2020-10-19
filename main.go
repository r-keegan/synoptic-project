package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/r-keegan/synoptic-project/Routes"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/r-keegan/synoptic-project/Models"
)

var DB *gorm.DB

func main() {
	database, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println("Failed to connect to database: ", err)
	}

	database.AutoMigrate(&Models.User{})
	DB = database

	r := Routes.SetupRouter()
	r.Run()
}
