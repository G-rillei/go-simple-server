package repository

import (
	"github.com/G-rillei/go-simple-server/model"
	"gorm.io/gorm"
)

type PostRepository struct {
	conn *gorm.DB
}

func NewPostRepository(conn *gorm.DB) PostRepository {
	return PostRepository{conn}
}

func (pr *PostRepository) GetAllPosts() ([]model.Post, error) {
	var posts []model.Post

	result := pr.conn.Find(&posts)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return make([]model.Post, 0), nil
		} else {
			return nil, result.Error
		}
	}

	return posts, nil
}

func (pr *PostRepository) CreatePost(post model.Post) (model.Post, error) {
	result := pr.conn.Create(&post)

	if result.Error != nil {
		return model.Post{}, result.Error
	}

	return post, nil
}
