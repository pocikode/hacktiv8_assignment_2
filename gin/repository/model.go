package repository

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item
}

type Item struct {
	gorm.Model
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint
}
