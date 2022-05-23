package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/z4vr/z4vr.dev/internal/config"
	"github.com/z4vr/z4vr.dev/internal/webserver"
)

var (
	flagConfig = flag.String("config", "config.toml", "Path to config file")
)

func main() {

	flag.Parse()

	// Create a new config
	cfg, err := config.Parse(*flagConfig, "z4vrdev_", config.DefaultConfig)
	if err != nil {
		logrus.WithError(err).Panic("Error loading config")
	}

	log := logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.TextFormatter{
			FullTimestamp:   true,
			ForceColors:     true,
			TimestampFormat: "2006-01-02 15:04:05",
		},
		Hooks: make(logrus.LevelHooks),
		Level: logrus.InfoLevel,
	}

	// Create a new webserver
	w := webserver.New(&cfg, &log)
	if w == nil {
		logrus.Panic("Error creating webserver")
	}

	// Listen
	if cfg.KeyFile != "" && cfg.CertFile != "" {
		if err := w.ListenTLS(); err != nil {
			logrus.WithError(err).Panic("Error listening")
		}
	} else {
		if err := w.Listen(); err != nil {
			logrus.WithError(err).Panic("Error listening")
		}
	}

	// block until exit
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}
