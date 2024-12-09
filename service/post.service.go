package service

import (
	"github.com/G-rillei/go-simple-server/model"
	"github.com/G-rillei/go-simple-server/repository"
)

type PostService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return PostService{repo}
}

func (ps *PostService) GetAllPosts() ([]model.Post, error) {
	return ps.repo.GetAllPosts()
}

func (ps *PostService) CreatePost(post model.Post) (model.Post, error) {
	return ps.repo.CreatePost(post)
}
