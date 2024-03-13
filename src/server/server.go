package server

import (
	"github.com/gofiber/fiber/v3"
)

func configureHandlers(app *fiber.App) {
	gp := app.Group("/clientes/:id")

	gp.Get("/extrato", extract)
	gp.Post("/transacoes", transaction)
}

func Run(port string) error {
	app := fiber.New()
	configureHandlers(app)

	return app.Listen("0.0.0.0:" + port)
}
