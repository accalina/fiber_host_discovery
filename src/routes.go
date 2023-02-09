package src

import (
	"github.com/accalina/fiber_host_discovery/src/api"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", api.GetHost)
	app.Get("/:name", api.GetHost)
	app.Post("/", api.CreateHost)
}
