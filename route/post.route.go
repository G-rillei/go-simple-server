package route

import (
	"github.com/G-rillei/go-simple-server/controller"
	"github.com/G-rillei/go-simple-server/repository"
	"github.com/G-rillei/go-simple-server/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PostRoutes(api fiber.Router, conn *gorm.DB) {
	router := api.Group("/posts")

	repository := repository.NewPostRepository(conn)
	service := service.NewPostService(repository)
	controller := controller.NewPostController(service)

	router.Get("/", controller.GetPosts)
}
