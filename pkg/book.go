package pkg

import (
	"encoding/json"
	"time"
)

type Book struct {
	ID          int
	Title       string
	Description string
	Price       json.Number
	Rating      json.Number
	Discount    json.Number
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BookResponse struct {
	ID          int         `json:"id"`
	Title       string      `json:"title" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Rating      json.Number `json:"rating" binding:"required,number"`
	Discount    json.Number `json:"discount" binding:"required,number"`
}
