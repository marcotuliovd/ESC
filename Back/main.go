package main

import (
	"Back/app/middleware"
	"Back/app/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    middleware.FiberMiddleware(app)

    routes.ServicoRoutes(app)
    routes.VagasRoutes(app)

    fmt.Println("ON in http://localhost:3000/")
}