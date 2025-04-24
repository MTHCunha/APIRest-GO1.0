package main

import (
	"github.com/MTHCunha/gocdd/configs"
	"github.com/MTHCunha/gocdd/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	routes.UserRoute(app)

	app.Listen(":8000")
}
