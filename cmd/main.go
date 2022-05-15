package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sarulabs/di"
	"github.com/z4vr/z4vr.dev/internal/config"
	"github.com/z4vr/z4vr.dev/internal/webserver"
)

var (
	flagConfigPath = flag.String("config", "./config.yaml", "Path to config file")
)

func main() {

	diBuilder, err := di.NewBuilder()
	if err != nil {
		panic(err)
	}

	diBuilder.Add(di.Def{
		Name: "config",
		Build: func(ctn di.Container) (interface{}, error) {
			p := config.NewConfitaProvider(*flagConfigPath)
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
		fmt.Sprintf(":%s", cfg.Instance().Port),
		cfg.Instance().CertFile, cfg.Instance().KeyFile)
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
