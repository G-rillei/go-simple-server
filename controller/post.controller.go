package controller

import (
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
