package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	// Cria uma nova aplicação Fiber
	app := fiber.New()

	// Rota para servir o vídeo
	app.Get("/video", func(c *fiber.Ctx) error {
		videoPath := "videos/movie.mp4"

		return c.SendFile(videoPath)
	})

	// Inicia o servidor na porta 8080
	log.Fatal(app.Listen(":8080"))
}
