package repository

import "gorm.io/gorm"

type PostRepository struct {
	conn *gorm.DB
}

func NewPostRepository(conn *gorm.DB) PostRepository {
	return PostRepository{conn}
}

func (pr *PostRepository) GetAllPosts() (string, error) {
	return "All posts", nil
}
