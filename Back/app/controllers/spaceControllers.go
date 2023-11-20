package controllers

import (
	"Back/Platform/database"
	"Back/app/models"
	"Back/app/utils"

	"github.com/gofiber/fiber/v2"
)

func EntryVehicle(c *fiber.Ctx) error {
	space := &models.Space{}

	if err := c.BodyParser(space); err != nil {
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

	if _, err := db.OccupationSpace(space); err != nil {
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
		"space":  space,
	})
}

func OutVehicle(c *fiber.Ctx) error {
	space := &models.OutVehicle{}

	if err := c.BodyParser(space); err != nil {
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
	occupation := &models.Space{}

	if err = db.GetVagaByVehicleId(space.VehicleId, occupation); err != nil {	// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	hourNow := utils.HourNow()

	timeOccupation := hourNow.Sub(occupation.Created_at)
    timeOccupationInt := int(timeOccupation.Minutes())

	amount := utils.CalculateAmount(occupation.Type, timeOccupationInt)

	history := &models.History{
		Amount: amount,
		VehicleId: occupation.VehicleId,
		Entry: occupation.Created_at,
		Exit: hourNow,
		Type: occupation.Type,
	}
	
	if err := db.CreateHistory(history); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := db.DeleteOccupation(occupation.Id); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	
	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"data":  history,
	})
}