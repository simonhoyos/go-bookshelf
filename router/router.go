package router

import (
	"go-bookshelf/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	author := api.Group("/authors")
	author.Get("/", controllers.GetAllAuthors)
	author.Get("/:id", controllers.GetAuthor)
	author.Post("/", controllers.CreateAuthor)
	author.Put("/:id", controllers.UpdateAuthor)
	author.Delete("/:id", controllers.DeleteAuthor)
}
