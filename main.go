package main

import (
	"context"
	"fmt"

	"github.com/appleboy/graceful"
	"github.com/caitlinelfring/go-env-default"
	"github.com/gofiber/fiber/v2"
)

var port = env.GetIntDefault("PORT", 3000)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Identity Service")
	})

	m := graceful.NewManager()

	m.AddRunningJob(func(ctx context.Context) error {
		return app.Listen(fmt.Sprintf(":%d", port))
	})

	m.AddShutdownJob(func() error {
		return app.Shutdown()
	})

	<-m.Done()
}
