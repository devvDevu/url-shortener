package app

import (
	"context"
	"url-shortener/cmd/app/adapters"
	"url-shortener/cmd/app/handlers"
	"url-shortener/cmd/app/repositories"
	"url-shortener/cmd/app/services"
	"url-shortener/cmd/app/usecases"
	"url-shortener/internal/config"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type App struct {
	cfg          *config.Config
	adapters     *adapters.Adapters
	usecases     *usecases.Usecases
	handlers     *handlers.Handlers
	services     *services.Services
	repositories *repositories.Repositories
}

func NewApp(cfg *config.Config) *App {
	return &App{
		cfg:          cfg,
		adapters:     adapters.NewAdapters(cfg),
		usecases:     usecases.NewUsecases(cfg),
		handlers:     handlers.NewHandlers(cfg),
		services:     services.NewServices(cfg),
		repositories: repositories.NewRepositories(cfg),
	}
}

func (a *App) GetAdapters() *adapters.Adapters {
	return a.adapters
}

func (a *App) MustInit(ctx context.Context, r *mux.Router) (*App, error) {
	const action = "App MustInit "
	var err error

	{
		logrus.Info(action, "adapters starting")
		err = a.adapters.MustInit(ctx)
		if err != nil {
			logrus.WithError(err).Error(action, "adapters failed")
			return nil, err
		}
		logrus.Info(action, "adapters running")
	}

	{
		logrus.Info(action, "repositories starting")
		err = a.repositories.MustInit(a.adapters)
		if err != nil {
			logrus.WithError(err).Error(action, "repositories failed")
			return nil, err
		}
		logrus.Info(action, "repositories running")
	}

	{
		logrus.Info(action, "services starting")
		err = a.services.MustInit(a.repositories)
		if err != nil {
			logrus.WithError(err).Error(action, "services failed")
			return nil, err
		}
		logrus.Info(action, "services running")
	}

	{
		logrus.Info(action, "usecases starting")
		err = a.usecases.MustInit(a.services)
		if err != nil {
			logrus.WithError(err).Error(action, "usecases failed")
			return nil, err
		}
		logrus.Info(action, "usecases running")
	}

	{
		logrus.Info(action, "handlers starting")
		logrus.Info(action, "----------------------------------------------------")
		err = a.handlers.MustInit(ctx, a.usecases, r)
		if err != nil {
			logrus.WithError(err).Error(action, "handlers failed")
			return nil, err
		}
		logrus.Info(action, "----------------------------------------------------")
		logrus.Info(action, "handlers running")
	}

	logrus.Info(action, "done")

	return a, nil
}
