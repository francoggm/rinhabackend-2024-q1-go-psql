package server

import (
	"net/http"
	"strconv"

	"github.com/francoggm/rinhabackend-2024-q1-go-psql/domain/client"
	"github.com/francoggm/rinhabackend-2024-q1-go-psql/infra/database"
	"github.com/gofiber/fiber/v3"
)

func extract(c fiber.Ctx, db *database.Database) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(http.StatusNotFound).Send(nil)
	}

	if id < 1 || id > 5 {
		return c.Status(http.StatusNotFound).Send(nil)
	}

	extract, err := db.GetExtract(c.Context(), id)
	if err != nil {
		switch err {
		case client.ErrNotFound:
			return c.Status(http.StatusNotFound).Send(nil)
		default:
			return c.Status(http.StatusUnprocessableEntity).Send(nil)
		}
	}

	return c.JSON(extract)
}

func transaction(c fiber.Ctx, db *database.Database) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(http.StatusNotFound).Send(nil)
	}

	if id < 1 || id > 5 {
		return c.Status(http.StatusNotFound).Send(nil)
	}

	var req client.TransactionReq
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).Send(nil)
	}

	if !req.IsValid() {
		return c.Status(http.StatusUnprocessableEntity).Send(nil)
	}

	res, err := db.MakeTransaction(c.Context(), id, req.Value, req.Description, req.Type)
	if err != nil {
		switch err {
		case client.ErrNotFound:
			return c.Status(http.StatusNotFound).Send(nil)
		case client.ErrInsufficientLimit:
			return c.Status(http.StatusUnprocessableEntity).Send(nil)
		default:
			return c.Status(http.StatusBadGateway).Send(nil)
		}
	}

	return c.JSON(res)
}
