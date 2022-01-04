package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Api for testing Stripe!")
	})

	app.Get("/SomethingCool", func(c *fiber.Ctx) error {
		data := make(map[string]string)
		data["FancyData"] = "Here we have all you need"
		return c.JSON(data)
	})

	app.Listen(":3000")
}
