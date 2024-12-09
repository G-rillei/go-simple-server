package repository

type PostRepository struct {
}

func NewPostRepository() PostRepository {
	return PostRepository{}
}

func (pr *PostRepository) GetAllPosts() (string, error) {
	return "All posts", nil
}
