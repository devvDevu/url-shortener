package services

import (
	"url-shortener/cmd/app/repositories"
	"url-shortener/internal/config"
	"url-shortener/internal/service/url_service"

	"github.com/sirupsen/logrus"
)

type Services struct {
	cfg *config.Config
	url *url_service.UrlService
}

func NewServices(cfg *config.Config) *Services {
	return &Services{cfg: cfg}
}

func (s *Services) GetUrl() *url_service.UrlService {
	return s.url
}

func (s *Services) MustInit(repositories *repositories.Repositories) error {
	const action = "Services MustInit"
	var err error

	{
		s.url = url_service.NewUrlService(repositories.GetUrl())

		logrus.WithFields(logrus.Fields{
			"serviceName": "UrlService",
		}).Info(action, " initialized")
	}

	logrus.Info(action, " done")

	return err
}
