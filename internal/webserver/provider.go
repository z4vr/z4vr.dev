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

	return prov

}
