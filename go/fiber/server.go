package main

import (
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	Id    int
	Name  string
	Price float32
}

func add(value1, value2 int) int {
	return value1 + value2
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	app.Get("/products", func(c *fiber.Ctx) error {
		products := []Product{
			{1, "Fridge1", 200.99},
			{2, "Fridge2", 300.99},
			{3, "Fridge3", 400.99},
			{4, "Fridge4", 500.99},
			{5, "Fridge5", 600.99},
		}
		return c.JSON(products)
	})

	app.Get("/add", func(c *fiber.Ctx) error {
		return c.JSON(add(3, 8))
	})

	// http://127.0.0.1:13004
	app.Listen(":13004")
}
