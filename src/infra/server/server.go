package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

func configureHandlers(app *fiber.App, db *pgxpool.Pool) {
	gp := app.Group("/clientes/:id")

	gp.Get("/extrato", func(c fiber.Ctx) error {
		return extract(c, db)
	})
	gp.Post("/transacoes", func(c fiber.Ctx) error {
		return transaction(c, db)
	})
}

func Run(port string, db *pgxpool.Pool) error {
	app := fiber.New()
	configureHandlers(app, db)

	return app.Listen("0.0.0.0:" + port)
}
