package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed views/*
var views embed.FS

func main() {
	app := fiber.New()

	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(views),
		PathPrefix: "views",
	}))

	// Start the Fiber server
	log.Fatal(app.Listen(":3000"))
}
