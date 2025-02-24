package config

import (
	"context"
	"os"
	"url-shortener/internal/common/types/error_with_codes"
	"url-shortener/internal/config/database"
	"url-shortener/internal/config/http"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Http      http.HttpConfig
	Database  database.DatabaseConfig
	Env       EnvConfig
	path      string
	envReader envReader
}

type envReader interface {
	EnvReadConfig(addr string, cfg interface{}) error
}

func MustLoad(ctx context.Context, configPath string, envReader envReader) *Config {
	operation := "config.MustLoad()"

	cfg := new(Config)
	cfg.envReader = envReader

	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		logrus.WithFields(logrus.Fields{
			"config_path": configPath,
		}).WithError(err).Fatal(error_with_codes.ErrorFailedToFindConfig.SetOperation(operation))
	}

	err = envReader.EnvReadConfig(configPath, cfg)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"config_path": configPath,
		}).WithError(err).Fatal(error_with_codes.ErrorFailedToReadConfig.SetOperation(operation))
	}

	return cfg
}
