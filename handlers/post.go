// handlers/postHandler.go
package handlers

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/models"
	"github.com/gofiber/fiber/v2"
)

// GetPosts gets all posts
func GetPosts(c *fiber.Ctx) error {
	var posts []models.Post
	database.PostDB.Find(&posts)
	return c.JSON(posts)
}

// CreatePost creates a new post
func CreatePost(c *fiber.Ctx) error {
	post := new(models.Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.PostDB.Create(&post)
	return c.Status(201).JSON(post)
}

// UpdatePost updates a post by ID
func UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	if err := database.PostDB.First(&post, id).Error; err != nil {
		return c.Status(404).SendString("Post not found")
	}
	if err := c.BodyParser(&post); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.PostDB.Save(&post)
	return c.JSON(post)
}

// DeletePost deletes a post by ID
func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	if err := database.PostDB.First(&post, id).Error; err != nil {
		return c.Status(404).SendString("Post not found")
	}
	database.PostDB.Delete(&post)
	return c.SendStatus(204)
}
