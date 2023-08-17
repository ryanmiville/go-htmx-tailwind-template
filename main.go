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
	public  embed.FS
	count   int
	devMode bool
)

func main() {
	if os.Getenv("DEV") == "true" {
		devMode = true
	}
	app := fiber.New(fiber.Config{
		Views: createEngine(),
	})
	initRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func createEngine() *django.Engine {
	engine := django.NewPathForwardingFileSystem(http.FS(views), "/views", ".html")
	if devMode {
		engine = django.New("./views", ".html")
	}
	engine.Reload(true)
	return engine
}

func initRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Count": count,
		})
	})

	if devMode {
		app.Static("/public", "./public")
	} else {
		app.Use("/public", filesystem.New(filesystem.Config{
			Root:       http.FS(public),
			PathPrefix: "public",
		}))
	}

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
