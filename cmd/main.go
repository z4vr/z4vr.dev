package main

import (
	"os"

	"github.com/z4vr/z4vr.dev/internal/routes"
)

const defaultPort string = "3000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := routes.SetupRoutes()

	r.Run(":" + port)
}
