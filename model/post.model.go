package model

import "time"

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey; autoIncrement"`
	Title     string    `json:"title" gorm:"not null"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreatePostModel struct {
	Title   string `json:"title" validate:"required,min=5,max=100"`
	Content string `json:"content"`
}
