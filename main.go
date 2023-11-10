package main

import (
	"os"

	"github.com/onainadapdap1/kartu_prakerja/simple_unjuk_ket/prakerja-final/routes"
	"github.com/onainadapdap1/kartu_prakerja/simple_unjuk_ket/prakerja-final/config"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
}

func main() {
	e := routes.Init()
	e.Logger.Fatal(e.Start(envPortOr("8000")))
}

func envPortOr(port string) string {
	envPort := os.Getenv("PORT")
	if envPort != "" {
		return ":" + envPort
	}
	return ":" + port
}
