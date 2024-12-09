package service

import "github.com/G-rillei/go-simple-server/repository"

type PostService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return PostService{repo}
}

func (ps *PostService) GetAllPosts() (string, error) {
	return ps.repo.GetAllPosts()
}
