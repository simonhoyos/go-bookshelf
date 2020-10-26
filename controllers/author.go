package controllers

import (
	"go-bookshelf/database"
	"go-bookshelf/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllAuthors(c *fiber.Ctx) error {
	db := database.DB

	var authors []models.Author
	db.Find(&authors)

	return c.JSON(fiber.Map{"status": "success", "message": "All authors", "data": authors})
}

func GetAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var author models.Author
	db.Find(&author, id)

	if author.Name == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No author found with ID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Author found", "data": author})
}

func CreateAuthor(c *fiber.Ctx) error {
	db := database.DB
	author := new(models.Author)

	if err := c.BodyParser(author); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	if err := db.Create(&author).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create author", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created author", "data": author})
}

func UpdateAuthor(c *fiber.Ctx) error {
	author := new(models.Author)

	if err := c.BodyParser(author); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	id := c.Params("id")
	db := database.DB

	var updateAuthor models.Author
	query := db.First(&updateAuthor, id)

	if updateAuthor.Name == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No author found with ID", "data": nil})
	}

	query.Updates(author)

	return c.JSON(fiber.Map{"status": "success", "message": "Author successfully updated", "data": updateAuthor})
}

func DeleteAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var author models.Author
	db.First(&author, id)

	if author.Name == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No author found with ID", "data": nil})
	}

	db.Delete(&author)

	return c.JSON(fiber.Map{"status": "success", "message": "Author successfully deleted", "data": author})
}
