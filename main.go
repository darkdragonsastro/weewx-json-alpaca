package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"

	"github.com/darkdragonsastro/weewx-json-alpaca/alpaca"
	"github.com/darkdragonsastro/weewx-json-alpaca/env"
	"github.com/darkdragonsastro/weewx-json-alpaca/handler"
	"github.com/darkdragonsastro/weewx-json-alpaca/logging"
	"github.com/darkdragonsastro/weewx-json-alpaca/router"
	"github.com/darkdragonsastro/weewx-json-alpaca/server"
	"github.com/darkdragonsastro/weewx-json-alpaca/weewx"
)

type config struct {
	env.Config
	ListenIPAddress string `env:"LISTEN_IP,required"`
	ListenPort      int    `env:"LISTEN_PORT,required"`
	WeeWxURL        string `env:"WEEWX_URL,required"`
}

func main() {
	var err error

	defer func() {
		// If we are existing due to an error, be sure to set the exit code appropriately.
		if err != nil {
			os.Exit(1)
		}
	}()

	var c config
	err = env.Load(&c)
	log := logging.Initialize(c.Config)

	if err != nil {
		log.Error("error initializing environment", zap.Error(err))
		return
	}

	log.Info("initializing")

	discovery := alpaca.NewAlpacaDiscovery(log, c.ListenPort)

	discovery.StartDiscovery()

	h := handler.New(weewx.NewClient(c.WeeWxURL, log))

	r := router.NewRouter(h, log)

	err = startHTTPServer(r, log, fmt.Sprintf("%s:%d", c.ListenIPAddress, c.ListenPort))
	if err != nil {
		log.Error("error starting server", zap.Error(err))
		return
	}
}

func startHTTPServer(h http.Handler, log *zap.Logger, serverAddr string) error {
	httpServer := &http.Server{
		Addr:    serverAddr,
		Handler: h,
	}

	ln, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Error("error starting listener", zap.Error(err))
		return err
	}

	log.Info("starting server", zap.String("addr", serverAddr))

	// Start the http server.
	s := server.NewGracefulHTTPServer(log, httpServer, ln, 30*time.Second)
	return s.Run()
}
