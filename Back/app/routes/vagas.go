package routes

import (
	"Back/app/controllers"

	"github.com/gofiber/fiber/v2"
)


func VagasRoutes(a *fiber.App) {

	route := a.Group("/api/v1")


	route.Post("/vagas/entrada", controllers.EntradaDeVeiculo)
	route.Post("/vagas/saida", controllers.SaidaDeVeiculo)
}
