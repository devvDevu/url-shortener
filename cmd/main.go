package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"url-shortener/cmd/app"
	custom "url-shortener/customize/rout"
	"url-shortener/internal/config"
	"url-shortener/internal/pkg/env"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const defaultConfigPath = "config/local.yaml"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	configPath := defaultConfigPath
	if os.Getenv("CONFIG_PATH") != "" {
		configPath = os.Getenv("CONFIG_PATH")
	}

	envReader := env.NewEnvReader()

	router := mux.NewRouter()

	cfg := config.MustLoad(ctx, configPath, envReader)

	app, err := app.NewApp(cfg).MustInit(ctx, router)
	if err != nil {
		logrus.WithError(err).Error("failed to init app")
		return
	}

	srv := &http.Server{
		Addr:    cfg.GetHttp().GetAddr(),
		Handler: router,
	}

	go func() {
		var srvErr error
		httpCfg := cfg.GetHttp()

		l := logrus.WithFields(logrus.Fields{
			"addr":     httpCfg.GetAddr(),
			"useHttps": httpCfg.GetUseHttps(),
		})

		l.Info("start http server")

		custom.PrintRoutesTable(router)

		if httpCfg.GetUseHttps() {
			//srvErr = srv.ListenAndServeTLS(httpCfg.GetCertFile().String(), httpCfg.GetKeyFile().String())
		} else {
			srvErr = srv.ListenAndServe()
		}
		if srvErr != nil && srvErr != http.ErrServerClosed {
			l.Fatalf("srv listen: %s\n", srvErr)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Wait for interrupt signal
	sig := <-sigChan
	logrus.Infof("received signal: %v", sig)

	// Create shutdown context with timeout
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	// Close database connections
	if err := app.GetAdapters().GetPostgres().Close(); err != nil {
		logrus.WithError(err).Error("failed to close database connection")
	}

	// Attempt graceful shutdown
	if err := srv.Shutdown(shutdownCtx); err != nil {
		logrus.WithError(err).Error("server shutdown failed")
	}

	logrus.Info("server shutdown completed")
}
