package url_get_handler

import (
	"context"
	"io"
	"net/http"
	"time"
	"url-shortener/internal/common/types/handler_type"
	"url-shortener/internal/common/types/url_types"
	"url-shortener/internal/data_transfer_object/result"
	url_get_dto_request "url-shortener/internal/data_transfer_object/url_dto/request/url_get"
	url_get_dto_response "url-shortener/internal/data_transfer_object/url_dto/response/url_get"

	"github.com/go-playground/validator"
	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
)

const handlerName = "UrlGetHandler"

type UrlGetHandler struct {
	usecase usecaseI
}

type usecaseI interface {
	GetUrlByCode(ctx context.Context, code url_types.UrlCode) (*url_types.UrlOriginal, error)
}

func NewUrlGetHandler(usecase usecaseI) *UrlGetHandler {
	return &UrlGetHandler{usecase: usecase}
}

func (h *UrlGetHandler) GetMethod() handler_type.HandlerMethod {
	return http.MethodGet
}

func (h *UrlGetHandler) GetPath() string {
	return "/url"
}

func (h *UrlGetHandler) ExecFunc(ctx context.Context, r *http.Request) ([]byte, error) {
	const action = "UrlGetHandler ExecFunc "
	const method = "ExecFunc"

	t := time.Now()

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"handlerName": handlerName,
			"method":      method,
		}).WithError(err).Error(action)
		return nil, err
	}

	var dtoRequest url_get_dto_request.UrlGetDto
	err = json.UnmarshalContext(ctx, body, &dtoRequest)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"handlerName": handlerName,
			"method":      method,
		}).WithError(err).Error(action)
		return nil, err
	}

	validate := validator.New()
	err = validate.Struct(dtoRequest)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"handlerName": handlerName,
			"method":      method,
		}).WithError(err).Error(action)
		return nil, err
	}
	originalUrl, err := h.usecase.GetUrlByCode(ctx, dtoRequest.Code)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"handlerName": handlerName,
			"method":      method,
		}).WithError(err).Error(action)
		return nil, err
	}

	dtoResponse := url_get_dto_response.UrlGetDto{
		OriginalUrl: *originalUrl,
	}

	json, err := result.NewResultOk(dtoResponse, time.Since(t)).GetJson()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"handlerName": handlerName,
			"method":      method,
		}).WithError(err).Error(action)
		return nil, err
	}

	return json, nil
}
