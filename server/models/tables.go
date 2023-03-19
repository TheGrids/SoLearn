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
	Role       byte       `json:"role" gorm:"default:1"`
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

type Course struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Likes       uint       `json:"likes"`
	Image       string     `json:"image"`
	Chapter     []Chapter  `json:"chapter"`
}

type Chapter struct {
	ID       uint      `gorm:"primary_key" json:"id"`
	CourseID uint      `json:"course_id"`
	Number   uint      `json:"number"`
	Name     string    `json:"name"`
	Subject  []Subject `json:"subject"`
}

type Subject struct {
	ChapterID uint    `json:"chapterID"`
	Number    float32 `json:"number"`
	Name      string  `json:"name"`
}
