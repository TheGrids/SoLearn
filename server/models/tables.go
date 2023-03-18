package models

import "time"

type User struct {
	ID         uint       `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	Email      string     `json:"email" binding:"required"`
	Password   string     `json:"password" binding:"required"`
	FirstName  string     `json:"first_name" binding:"required"`
	LastName   string     `json:"last_name" binding:"required"`
	Role       Role       `json:"role" gorm:"default:1"`
	EmailCheck EmailCheck `json:"email_check"`
	Token      []Token    `json:"token"`
}

type EmailCheck struct {
	UserID uint   `json:"user_id"`
	UUID   string `json:"uuid"`
}

type Token struct {
	UserID  uint   `json:"user-id"`
	Refresh string `json:"refresh"`
}
