package controller

import (
	"github.com/G-rillei/go-simple-server/model"
	"github.com/G-rillei/go-simple-server/service"
	"github.com/gofiber/fiber/v2"
)

type PostController struct {
	service service.PostService
}

func NewPostController(service service.PostService) PostController {
	return PostController{service}
}

func (pc *PostController) GetPosts(c *fiber.Ctx) error {
	posts, err := pc.service.GetAllPosts()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"posts": posts})
}

func (pc *PostController) CreatePost(c *fiber.Ctx) error {
	var body model.CreatePostModel

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	post := model.Post{
		Title:   body.Title,
		Content: body.Content,
	}

	createdPost, err := pc.service.CreatePost(post)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(createdPost)
}
