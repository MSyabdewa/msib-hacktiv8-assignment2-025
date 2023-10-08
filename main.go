package main

import (
	"github.com/MSyabdewa/msib-hacktiv8-assignment2-025/database"
	"github.com/MSyabdewa/msib-hacktiv8-assignment2-025/routes"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	var err error
	db, err = database.InitializeDatabase()
	if err != nil {
		panic("Failed to connect to database")
	}

	r := routes.SetupRouter()
	r.Run(":8080")
}
