package repositories

import (
	"url-shortener/cmd/app/adapters"
	"url-shortener/internal/config"
	"url-shortener/internal/repository/url_repository"

	"github.com/sirupsen/logrus"
)

type Repositories struct {
	cfg *config.Config
	url *url_repository.UrlRepository
}

func NewRepositories(cfg *config.Config) *Repositories {
	r := new(Repositories)
	r.cfg = cfg

	return r
}

func (r *Repositories) GetUrl() *url_repository.UrlRepository {
	return r.url
}

func (r *Repositories) MustInit(adapters *adapters.Adapters) error {
	const action = "Repositories MustInit"
	var err error

	{
		r.url = url_repository.NewUrlRepository(adapters.GetPostgres())

		logrus.WithFields(logrus.Fields{
			"repositoryName": "UrlRepository",
		}).Info(action, " initialized")
	}

	logrus.Info(action, " done")

	return err
}
