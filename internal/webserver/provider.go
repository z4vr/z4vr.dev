package webserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sarulabs/di"
	"github.com/z4vr/z4vr.dev/internal/config"
)

type Provider struct {
	App *fiber.App
}

func NewFiberProvider(ctn di.Container) *Provider {

	cfg := ctn.Get("config").(config.Provider)

	prov := &Provider{
		App: fiber.New(),
	}

	prov.App.Static("/", cfg.Instance().StaticDir)

	// register the api group
	apiGroup := prov.App.Group("/api", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
			"message": "Your request was processed successfully, but nothing is here.",
		})
	})

	// register the api routes
	apiGroup.Get("/gear", func(c *fiber.Ctx) error {
		// return the gear list
		return c.Status(fiber.StatusOK).JSON(nil)
	})

	return prov

}
