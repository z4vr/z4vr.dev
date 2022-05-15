package webserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/sarulabs/di"
	"github.com/z4vr/z4vr.dev/internal/config"
)

type Provider struct {
	App *fiber.App
}

func NewFiberProvider(ctn di.Container) *Provider {

	cfg := ctn.Get("config").(config.Provider)

	prov := &Provider{
		App: fiber.New(fiber.Config{
			AppName: "z4vr.dev",
		}),
	}

	prov.App.Use(favicon.New(favicon.Config{
		File: cfg.Instance().Fiber.StaticDir + "/favicon.ico",
	}))

	prov.App.Static("/", cfg.Instance().Fiber.StaticDir)

	return prov

}
