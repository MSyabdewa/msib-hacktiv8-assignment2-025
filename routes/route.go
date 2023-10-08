package routes

import (
	"github.com/MSyabdewa/msib-hacktiv8-assignment2-025/database"
	handler "github.com/MSyabdewa/msib-hacktiv8-assignment2-025/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Inisialisasi koneksi database PostgreSQL
	db, err := database.InitializeDatabase()
	if err != nil {
		panic("Failed to connect to database")
	}

	// Mengirimkan koneksi database sebagai middleware ke handler
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Definisi route
	r.POST("/orders", handler.CreateNewOrder)
	r.PUT("/orders/:order_id", handler.UpdateOrder)
	r.PATCH("/orders/:order_id", handler.UpdateOrder)
	r.GET("/orders", handler.GetAllOrder)
	r.DELETE("/orders/:order_id", handler.DeleteOrder)

	return r
}
