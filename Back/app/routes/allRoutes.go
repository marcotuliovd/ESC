package routes

import (
	"Back/app/controllers"

	"github.com/gofiber/fiber/v2"
)


func AllRoutes(a *fiber.App) {

	route := a.Group("/api/v1")
	route.Post("/createUser", controllers.CreateUser)
	route.Post("/createVehicle", controllers.CreateVehicle)
	route.Post("/occupationSpace", controllers.EntryVehicle)
}
