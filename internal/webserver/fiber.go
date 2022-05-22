package webserver

import (
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

	w.SetupRoutes()

	return w

}

func (w *Webserver) Listen() error {
	return w.app.Listen(w.config.Port)
}

func (w *Webserver) ListenTLS() error {
	return w.app.ListenTLS(w.config.Port, w.config.CertFile, w.config.KeyFile)
}

func (w *Webserver) SetupRoutes() {
	w.app.Static("/", w.config.StaticDir)
}