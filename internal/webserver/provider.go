package webserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sarulabs/di"
)

type Provider struct {
	App *fiber.App
}

func NewFiberProvider(ctn di.Container) *Provider {

	return &Provider{
		App: fiber.New(),
	}
}
