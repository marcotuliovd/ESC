package controllers

import (
	"Back/Platform/database"
	"Back/app/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func EntradaDeVeiculo(c *fiber.Ctx) error {
	vaga := &models.Vaga{}

	if err := c.BodyParser(vaga); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	fmt.Println(vaga)

	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := db.EntradaDeVeiculo(vaga); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"vaga":  vaga,
	})
}