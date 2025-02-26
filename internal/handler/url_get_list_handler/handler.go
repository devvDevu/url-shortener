package url_get_list_handler

import (
	"context"
	"net/http"
	"time"
	"url-shortener/internal/common/types/handler_type"
	"url-shortener/internal/common/types/url_types"
	"url-shortener/internal/data_transfer_object/result"
	url_get_list_dto "url-shortener/internal/data_transfer_object/url_dto/response/url_get_list"
	"url-shortener/internal/value_object/url_value_object"

	"github.com/sirupsen/logrus"
)

const handlerName = "UrlGetListHandler"

type UrlGetListHandler struct {
	usecase usecaseI
}

type usecaseI interface {
	GetUrlList(ctx context.Context) (map[url_types.UrlId]*url_value_object.Url, error)
}

func NewUrlGetListHandler(usecase usecaseI) *UrlGetListHandler {
	return &UrlGetListHandler{usecase: usecase}
}

func (h *UrlGetListHandler) GetMethod() handler_type.HandlerMethod {
	return http.MethodGet
}

func (h *UrlGetListHandler) GetPath() handler_type.HandlerPath {
	return "/url/list"
}

func (h *UrlGetListHandler) ExecFunc(ctx context.Context, r *http.Request) ([]byte, error) {
	const action = "UrlGetListHandler ExecFunc "
	const method = "ExecFunc"

	t := time.Now()

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	urlList, err := h.usecase.GetUrlList(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"handlerName": handlerName,
			"method":      method,
		}).WithError(err).Error(action)
		return nil, err
	}
	logrus.Info(urlList)

	dtoResponse := url_get_list_dto.NewUrlGetListDto(urlList)

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
