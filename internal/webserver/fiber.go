package webserver

import (
	"fmt"
	"path"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/sirupsen/logrus"

	"github.com/z4vr/z4vr.dev/internal/config"
)

type Webserver struct {
	config *config.Config
	logger logrus.FieldLogger

	app *fiber.App
}

func New(config *config.Config, logger logrus.FieldLogger) *Webserver {
	w := &Webserver{
		config: config,
		logger: logger,

		app: fiber.New(),
	}

	w.logger.WithFields(logrus.Fields{
		"addr": w.config.Address,
		"port": config.Port,
	}).Info("Starting webserver")

	w.app.Use(favicon.New(favicon.Config{
		File: path.Join(w.config.StaticDir, "favicon.ico"),
	}))

	w.app.Use(logrus.New())

	w.SetupRoutes()

	return w

}

func (w *Webserver) Listen() error {
	return w.app.Listen(fmt.Sprintf("%s:%s", w.config.Address, w.config.Port))
}

func (w *Webserver) ListenTLS() error {
	return w.app.ListenTLS(fmt.Sprintf("%s:%s", w.config.Address, w.config.Port), w.config.CertFile, w.config.KeyFile)
}

func (w *Webserver) SetupRoutes() {
	w.app.Static("/", w.config.StaticDir, fiber.Static{
		Index: w.config.IndexFile,
	})
}
