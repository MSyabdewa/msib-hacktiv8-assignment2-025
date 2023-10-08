package handler

import (
	"fmt"
	"net/http"

	"github.com/MSyabdewa/msib-hacktiv8-assignment2-025/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateNewOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, _ := c.Get("db")
	database := db.(*gorm.DB)

	// Check if item_code already exists in the database
	var existingItem models.Item
	if err := database.Where("item_code = ?", order.Items[0].ItemCode).First(&existingItem).Error; err == nil {
		// If item_code already exists, send error response
		errorMessage := "Item code must be unique"
		fmt.Println(errorMessage)
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	// Create a new order in the database if item_code does not exist
	database.Create(&order)

	fmt.Println("Order created with id:", order.OrderID)
	c.JSON(http.StatusCreated, order)
}

func UpdateOrder(c *gin.Context) {
	// Get order_id from URL parameter
	orderID := c.Param("order_id")

	// Find order in the database based on order_id
	var order models.Order
	db, _ := c.Get("db")
	database := db.(*gorm.DB)
	if err := database.Where("order_id = ?", orderID).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Parse JSON data from the request
	var updateData models.UpdateOrderRequest
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use GORM for partial attribute updates on the order
	if err := database.Model(&order).Updates(map[string]interface{}{
		"ordered_at":    updateData.OrderedAt,
		"customer_name": updateData.CustomerName,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	// To update items, iterate through updateData.Items
	for _, itemData := range updateData.Items {
		// Find item in the database based on item_code
		var item models.Item
		if err := database.Where("item_code = ?", itemData.ItemCode).First(&item).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}

		// Use GORM to update the item
		if err := database.Model(&item).Updates(map[string]interface{}{
			"description": itemData.Description,
			"quantity":    itemData.Quantity,
		}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
			return
		}
	}

	fmt.Println("Order and Item successfully updated")
	c.JSON(http.StatusOK, gin.H{"message": "Order and Item successfully updated"})
}

func GetAllOrder(c *gin.Context) {
	db, _ := c.Get("db")
	database := db.(*gorm.DB)

	var orders []models.Order
	if err := database.Preload("Items").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order data"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func DeleteOrder(c *gin.Context) {
	// Get order ID from the parameter.
	orderID := c.Param("id")

	// Find order based on ID.
	db, _ := c.Get("db")
	database := db.(*gorm.DB)

	var order models.Order
	if err := database.Preload("Items").First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Delete items related to the order.
	if err := database.Delete(&order.Items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Delete order from the database.
	if err := database.Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send success response.
	fmt.Println("Order and Item successfully deleted")
	c.JSON(http.StatusOK, gin.H{"message": "Order and Item successfully deleted"})
}
