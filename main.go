package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Correct struct with proper backtick quotes
type AddOperationRequest struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

type AddOperationResponse struct {
	Result int `json:"result"`
}

func main() {
	fmt.Println("Starting server at http://localhost:3000")

	app := fiber.New()

	app.Get("/abc", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! 1234")
	})

	app.Post("/sum", AddOperation)

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func AddOperation(c *fiber.Ctx) error {
	var request AddOperationRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result := request.Num1 + request.Num2
	response := AddOperationResponse{Result: result}
	return c.JSON(response)
}

