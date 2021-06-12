package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
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
	// Echo instance
	app := echo.New()

	app.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello world")
	})

	app.GET("/products", func(context echo.Context) error {
		products := []Product{
			{1, "Fridge1", 200.99},
			{2, "Fridge2", 300.99},
			{3, "Fridge3", 400.99},
			{4, "Fridge4", 500.99},
			{5, "Fridge5", 600.99},
		}
		return context.JSON(http.StatusOK, products)
	})

	app.GET("/add", func(context echo.Context) error {
		return context.JSON(http.StatusOK, add(3, 8))
	})

	// http://127.0.0.1:13006
	app.Logger.Fatal(app.Start(":13006"))
}
