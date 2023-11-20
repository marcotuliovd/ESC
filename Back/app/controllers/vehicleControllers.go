package controllers

import (
	"Back/Platform/database"
	"Back/app/models"

	"github.com/gofiber/fiber/v2"
)

func CreateVehicle(c *fiber.Ctx) error {
	vehicle := &models.Vehicle{}

	if err := c.BodyParser(vehicle); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if _, err := db.CreateVehicle(vehicle); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

		return c.JSON(fiber.Map{
		"error": false,
		"id":   vehicle.Id,
		"vaga":  vehicle,
	})
}