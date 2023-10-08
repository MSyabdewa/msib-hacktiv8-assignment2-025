package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	OrderID      int `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	CustomerName string         `json:"customer_name"`
	OrderedAt    time.Time      `json:"ordered_at"`
	Items        []Item
}

type Item struct {
	ItemID      int    `gorm:"primaryKey"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type UpdateOrderRequest struct {
	OrderedAt    string              `json:"ordered_at"`
	CustomerName string              `json:"customer_name"`
	Items        []UpdateItemRequest `json:"items"`
}

type UpdateItemRequest struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
