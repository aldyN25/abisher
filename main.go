package main

import (
	"io"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/cinco/pkg/postgres"
	"gitlab.com/cinco/routes"
)

func main() {

	app := fiber.New(fiber.Config{})

	db := postgres.ConnectDB()

	routes.AllRouter(app, db)
	app.Listen(":8000")
}

type Views interface {
	Load() error
	Render(io.Writer, string, interface{}, ...string) error
}
