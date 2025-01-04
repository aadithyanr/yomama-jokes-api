package main

import (
    "log"
    "os"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    // da jokes
    app.Get("/jokes", func(c *fiber.Ctx) error {
        file, err := os.ReadFile("./jokes/jokes.txt")
        if err != nil {
            log.Printf("Error reading jokes file: %v", err)
            return c.Status(fiber.StatusInternalServerError).SendString("Could not load jokes")
        }
        c.Set("Content-Type", "text/plain")
        return c.Send(file)
    })

    // health endpoint
    app.Get("/ping", func(c *fiber.Ctx) error {
        c.Set("Content-Type", "text/plain")
        return c.SendString("pong")
    })

    log.Println("Server running on http://localhost:8000")
    log.Fatal(app.Listen(":8000"))
}
