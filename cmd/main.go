package main

import (
	"github.com/z4vr/z4vr.dev/internal/config"
	"github.com/z4vr/z4vr.dev/internal/routes"
)

var (
	cfgProvider = config.NewConfitaProvider()
)

func main() {

	cfg := getCfg()

	r := routes.SetupRoutes()

	r.Run(":" + cfg.Port)
}

func getCfg() *config.Config {
	cfgProvider.Load()
	return cfgProvider.Instance()
}
