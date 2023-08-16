package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/django/v3"
)

//go:embed views/*
var views embed.FS

//go:embed public/*
var public embed.FS

var count int

func main() {
	engine := django.NewPathForwardingFileSystem(http.FS(views), "/views", ".html")
	engine.Reload(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Count": count,
		})
	})

	app.Use("/public", filesystem.New(filesystem.Config{
		Root:       http.FS(public),
		PathPrefix: "public",
	}))

	app.Post("/increase", func(c *fiber.Ctx) error {
		count = count + 1
		return c.Render("partials/counter", fiber.Map{
			"Count": count,
		})
	})

	app.Post("/decrease", func(c *fiber.Ctx) error {
		count = count - 1
		return c.Render("partials/counter", fiber.Map{
			"Count": count,
		})
	})

	// Start the Fiber server
	log.Fatal(app.Listen(":3000"))
}
