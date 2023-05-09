package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mod/internal/auth/handlers"
	"log"
)

func main() {
	fmt.Println("hi")

	app := fiber.New()
	handlers.SetupRoutes(app)
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
