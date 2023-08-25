package routes

import (
	"github.com/gofiber/fiber/v2"
)


func VagasRoutes(a *fiber.App) {

	route := a.Group("/api/v1")


	route.Get("/books")
	route.Get("/book/:id")         
	route.Get("/token/new") 
}
