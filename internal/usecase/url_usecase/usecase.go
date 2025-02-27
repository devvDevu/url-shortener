package url_usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"strings"
	"url-shortener/internal/common/types/url_types"
	"url-shortener/internal/model/url_model"
	"url-shortener/internal/value_object/url_value_object"

	"github.com/sirupsen/logrus"
)

const (
	usecaseName = "UrlUsecase"
	codeLength  = 8
)

type UrlUsecase struct {
	service serviceI
}

type serviceI interface {
	GetUrlByCode(ctx context.Context, code url_types.UrlCode) (*url_model.Url, bool, error)
	CreateUrl(ctx context.Context, url *url_model.Url) error
	GetUrlList(ctx context.Context) ([]*url_model.Url, bool, error)
}

func NewUrlUsecase(service serviceI) *UrlUsecase {
	return &UrlUsecase{service: service}
}

func (u *UrlUsecase) MakeShortUrl(ctx context.Context, originalUrl url_types.UrlOriginal) (bool, error) {
	const action = "UrlUsecase MakeShortUrl "
	const method = "MakeShortUrl"

	code := url_types.UrlCode(generateSecureCode(codeLength))
	urlModel := &url_model.Url{
		Original: originalUrl,
		Code:     code,
	}

	err := u.service.CreateUrl(ctx, urlModel)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"usecaseName": usecaseName,
			"method":      method,
		}).WithError(err).Error(action)
		return false, err
	}

	return true, nil
}

func (u *UrlUsecase) GetUrlList(ctx context.Context) (map[url_types.UrlId]*url_value_object.Url, error) {
	const action = "UrlUsecase GetUrlList "
	const method = "GetUrlList"

	urls, ok, err := u.service.GetUrlList(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"usecaseName": usecaseName,
			"method":      method,
		}).WithError(err).Error(action)
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	urlsMap := make(map[url_types.UrlId]*url_value_object.Url)
	for _, url := range urls {
		urlsMap[url.Id] = &url_value_object.Url{
			Original: url.Original,
			Code:     url.Code,
		}
	}

	return urlsMap, nil
}

func (u *UrlUsecase) GetUrlByCode(ctx context.Context, code url_types.UrlCode) (*url_types.UrlOriginal, error) {
	const action = "UrlUsecase GetUrlByCode "
	const method = "GetUrlByCode"

	url, ok, err := u.service.GetUrlByCode(ctx, code)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"usecaseName": usecaseName,
			"method":      method,
		}).WithError(err).Error(action)
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	return &url.Original, nil
}

func generateSecureCode(length int) string {
	randomBytes := make([]byte, length)
	rand.Read(randomBytes)
	code := base64.URLEncoding.EncodeToString(randomBytes)
	code = strings.ReplaceAll(code, "/", "")
	code = strings.ReplaceAll(code, "+", "")
	code = strings.ReplaceAll(code, "=", "")
	return code[:length]
}
