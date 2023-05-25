package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sahublr01/rest_crud_temp/config"
	"github.com/sahublr01/rest_crud_temp/server"

	"github.com/kataras/golog"
)

func main() {
	if !config.Parse("") {
		golog.Error("Invalid Config provided")
		return
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signals
		golog.Info("crud_rest_temp stopped")

		os.Exit(0)
	}()

	golog.Info("Starting app...")
	server.Init()
	golog.Info("Exiting app...")

}
