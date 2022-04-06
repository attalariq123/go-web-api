package pkg

import "time"

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	Address   string `json:"address" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Address  string `json:"address" binding:"required"`
}
