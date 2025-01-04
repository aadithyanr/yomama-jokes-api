package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// da joke
	app.Get("/jokes", func(c *fiber.Ctx) error {
		file, err := os.ReadFile("./jokes/jokes.txt")
		if err != nil {
			log.Printf("Error reading jokes file: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Could not load jokes")
		}

		jokes := strings.Split(string(file), "\n")
		if len(jokes) == 0 {
			return c.Status(fiber.StatusInternalServerError).SendString("No jokes found")
		}

		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(jokes))

		// one joke for y'all
		c.Set("Content-Type", "text/plain")
		return c.SendString(jokes[randomIndex])
	})

	// health endpoint
	app.Get("/ping", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/plain")
		return c.SendString("pong")
	})

	log.Println("Server running on http://localhost:8000")
	log.Fatal(app.Listen(":8000"))
}
