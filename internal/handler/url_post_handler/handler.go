package url_post_handler

import (
	"context"
	"io"
	"net/http"
	"time"
	"url-shortener/internal/common/types/error_with_codes"
	"url-shortener/internal/common/types/handler_type"
	"url-shortener/internal/common/types/url_types"
	"url-shortener/internal/data_transfer_object/result"
	"url-shortener/internal/data_transfer_object/url_dto/request/request_url_post"
	"url-shortener/internal/data_transfer_object/url_dto/response/response_url_post"

	"github.com/go-playground/validator"
	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
)

const handlerName = "UrlPostHandler"

type UrlPostHandler struct {
	usecase usecaseI
}

type usecaseI interface {
	MakeShortUrl(ctx context.Context, originalUrl url_types.UrlOriginal) (bool, error)
}

func NewUrlPostHandler(usecase usecaseI) *UrlPostHandler {
	return &UrlPostHandler{usecase: usecase}
}

func (h *UrlPostHandler) GetMethod() handler_type.HandlerMethod {
	return http.MethodPost
}

func (h *UrlPostHandler) GetPath() handler_type.HandlerPath {
	return "/url"
}

// @Title Create short URL
// @Description Convert long URL to short code
// @Tags URL
// @Accept  json
// @Produce  json
// @Param   body body request_url_post.UrlPostDto true "URL Data"
// @Success 200 {object} response_url_post.UrlPostDto
// @Failure 400 {object} result.ResultErr
// @Failure 500 {object} result.ResultErr
// @Router /url [post]
func (h *UrlPostHandler) ExecFunc(ctx context.Context, r *http.Request) ([]byte, error) {
	const action = "UrlPostHandler ExecFunc "
	const method = "ExecFunc"

	t := time.Now()

	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"handlerName": handlerName,
			"method":      method,
		}).WithError(err).Error(action)
		return nil, err
	}

	var dtoRequest request_url_post.UrlPostDto
	logrus.Info(dtoRequest)
	err = json.UnmarshalContext(ctx, body, &dtoRequest)
	if err != nil {
		err := error_with_codes.ErrorFailedToCast
		logrus.WithFields(logrus.Fields{
			"handlerName": handlerName,
			"method":      method,
		}).WithError(err).Error(action)
		return nil, err
	}

	validate := validator.New()
	err = validate.Struct(dtoRequest)
	if err != nil {
		err := error_with_codes.ErrorFailedToValidate
		logrus.WithFields(logrus.Fields{
			"handlerName": handlerName,
			"method":      method,
		}).WithError(err).Error(action)
		return nil, err
	}

	ok, err := h.usecase.MakeShortUrl(ctx, dtoRequest.OriginalUrl)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"handlerName": handlerName,
			"method":      method,
		}).WithError(err).Error(action)
		return nil, err
	}

	dtoResponse := response_url_post.UrlPostDto{
		Ok: ok,
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
