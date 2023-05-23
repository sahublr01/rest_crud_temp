package main

import (
	"os"
	"os/signal"
	"syscall"

	"rest_crud_temp/config"
	"rest_crud_temp/server"

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
		// config.NotifyEvent("starting upgrade m3u8")
		<-signals
		golog.Info("upgrade m3u8 stopped")
		// config.NotifyEvent("stopping upgrade m3u8")
		os.Exit(0)
	}()

	golog.Info("Starting app...")
	server.Init()
	golog.Info("Exiting app...")

}
