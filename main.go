package main

import (
	"github.com/accalina/fiber_host_discovery/src"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	src.Routes(app)

	app.Listen(":8080")
}
