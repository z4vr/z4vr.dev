package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sarulabs/di"
	"github.com/z4vr/z4vr.dev/internal/config"
	"github.com/z4vr/z4vr.dev/internal/webserver"
)

func main() {

	loadenv("config.env")

	diBuilder, err := di.NewBuilder()
	if err != nil {
		panic(err)
	}

	diBuilder.Add(di.Def{
		Name: "config",
		Build: func(ctn di.Container) (interface{}, error) {
			p := config.NewConfitaProvider()
			return p, p.Load()
		},
	})

	diBuilder.Add(di.Def{
		Name: "webserver",
		Build: func(ctn di.Container) (interface{}, error) {
			return webserver.NewFiberProvider(ctn), nil
		},
	})

	ctn := diBuilder.Build()
	cfg := ctn.Get("config").(config.Provider)
	webserver := ctn.Get("webserver").(*webserver.Provider)
	err = webserver.App.ListenTLS(
		fmt.Sprintf("%s:%s", cfg.Instance().Fiber.Address, cfg.Instance().Fiber.Port),
		cfg.Instance().Fiber.CertFile, cfg.Instance().Fiber.KeyFile)
	if err != nil {
		log.Fatal(err)
	}

	block()

	err = ctn.DeleteWithSubContainers()
	if err != nil {
		return
	}

}

func block() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func loadenv(optionalFiles ...string) {
	for _, f := range optionalFiles {
		godotenv.Load(f)
	}
}
