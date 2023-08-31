package main

import (
	"Back/app/middleware"
	"Back/app/routes"
	"Back/app/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    middleware.FiberMiddleware(app)

    routes.VagasRoutes(app)

    utils.StartServer(app)

    fmt.Println("ON in http://localhost:3000/")
}