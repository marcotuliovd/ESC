package main

import (
	"Back/app/config"
	"Back/app/middleware"
	"Back/app/routes"
	"Back/app/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnvironmentVariables()
    app := fiber.New()

    middleware.FiberMiddleware(app)

    routes.AllRoutes(app)

    utils.StartServer(app)

    fmt.Println("ON in http://localhost:3000/")
}