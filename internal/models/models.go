package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type Poll struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	OwnerID     uint      `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

type Question struct {
	ID           uint      `gorm:"primaryKey"`
	PollID       uint      `gorm:"not null"`
	Text         string    `gorm:"not null"`
	QuestionType string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}

type Answer struct {
	ID         uint      `gorm:"primaryKey"`
	QuestionID uint      `gorm:"not null"`
	Text       string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

type Response struct {
	ID        uint      `gorm:"primaryKey"`
	PollID    uint      `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type ResponseAnswer struct {
	ID         uint `gorm:"primaryKey"`
	ResponseID uint `gorm:"not null"`
	QuestionID uint `gorm:"not null"`
	AnswerID   uint `gorm:"not null"`
}
