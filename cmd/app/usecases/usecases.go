package usecases

import (
	"url-shortener/cmd/app/services"
	"url-shortener/internal/config"
	"url-shortener/internal/usecase/url_usecase"

	"github.com/sirupsen/logrus"
)

type Usecases struct {
	cfg *config.Config
	url *url_usecase.UrlUsecase
}

func NewUsecases(cfg *config.Config) *Usecases {
	return &Usecases{cfg: cfg}
}

func (u *Usecases) GetUrl() *url_usecase.UrlUsecase {
	return u.url
}

func (u *Usecases) MustInit(services *services.Services) error {
	const action = "Usecases MustInit"
	var err error

	{
		u.url = url_usecase.NewUrlUsecase(services.GetUrl())

		logrus.WithFields(logrus.Fields{
			"usecaseName": "UrlUsecase",
		}).Info(action, " initialized")
	}

	logrus.Info(action, " done")

	return err
}
