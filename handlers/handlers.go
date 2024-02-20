// @/handlers/user.go
package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Getuser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user entities.user

	result := config.Database.Find(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&user)
}

func Adduser(c *fiber.Ctx) error {
	user := new(entities.user)

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&user)
	return c.Status(201).JSON(user)
}

func Updateuser(c *fiber.Ctx) error {
	user := new(entities.user)
	id := c.Params("id")

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&user)
	return c.Status(200).JSON(user)
}

func Removeuser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user entities.user

	result := config.Database.Delete(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
