package adapters

import (
	"context"
	"url-shortener/internal/adapter/database"
	"url-shortener/internal/config"

	"github.com/sirupsen/logrus"
)

type Adapters struct {
	cfg      *config.Config
	postgres *database.PostgresAdapter
}

func NewAdapters(cfg *config.Config) *Adapters {
	a := new(Adapters)
	a.cfg = cfg

	return a
}

// GetPostgres - getter for postgres adapter
func (a *Adapters) GetPostgres() *database.PostgresAdapter {
	return a.postgres
}

func (a *Adapters) MustInit(ctx context.Context) error {
	const action = "Adapters MustInit"
	var err error

	{
		a.postgres, err = database.NewPostgresAdapter(ctx, a.cfg.GetDatabase().GetPostgres())
		if err != nil {
			logrus.Fatalf("%s: %s: %s", action, "postgres", err.Error())
		}

		logrus.WithFields(logrus.Fields{
			"adapterName": "NewPostgresAdapter",
		}).Info(action, " initialized")
	}

	logrus.Info(action, " done")
	return err
}
