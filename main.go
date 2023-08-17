package main

import (
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/django/v3"
)

var (
	//go:embed views/*
	views embed.FS
	//go:embed public/*
	public embed.FS
	count  int
)

func main() {
	app := createApp()
	if os.Getenv("DEV") == "true" {
		app = createAppDev()
	}

	initRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

// createApp creates a fiber app with embedded views and assets
func createApp() *fiber.App {
	engine := django.NewPathForwardingFileSystem(http.FS(views), "/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use("/public", filesystem.New(filesystem.Config{
		Root:       http.FS(public),
		PathPrefix: "public",
	}))
	return app
}

// createAppDev creates a fiber app that reads views and assets from disk
func createAppDev() *fiber.App {
	engine := django.New("./views", ".html")
	engine.Reload(true)
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/public", "./public")
	return app
}

func initRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Count": count,
		})
	})

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
}
