package main

import (
	"APIREST-GO1.0/configs" //add this

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	app.Listen(":8000")
}
