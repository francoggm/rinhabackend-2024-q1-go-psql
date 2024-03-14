package server

import (
	"github.com/francoggm/rinhabackend-2024-q1-go-psql/infra/database"
	"github.com/gofiber/fiber/v3"
)

func configureHandlers(app *fiber.App, db *database.Database) {
	gp := app.Group("/clientes/:id")

	gp.Get("/extrato", func(c fiber.Ctx) error {
		return extract(c, db)
	})
	gp.Post("/transacoes", func(c fiber.Ctx) error {
		return transaction(c, db)
	})
}

func Run(port string, db *database.Database) error {
	app := fiber.New()
	configureHandlers(app, db)

	return app.Listen("0.0.0.0:" + port)
}
