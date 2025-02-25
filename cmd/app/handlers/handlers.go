package handlers

import (
	"context"
	"net/http"
	"time"
	"url-shortener/cmd/app/usecases"
	"url-shortener/internal/common/types/error_with_codes"
	"url-shortener/internal/common/types/handler_type"
	"url-shortener/internal/config"
	"url-shortener/internal/data_transfer_object/result"
	"url-shortener/internal/handler/url_get_handler"
	"url-shortener/internal/handler/url_get_list_handler"
	"url-shortener/internal/handler/url_post_handler"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const defaultExecTimeout = 3 * time.Second

type Handlers struct {
	cfg *config.Config
}

func NewHandlers(cfg *config.Config) *Handlers {
	return &Handlers{cfg}
}

func (h *Handlers) MustInit(ctx context.Context, usecases *usecases.Usecases, router *mux.Router) error {
	const action = "Handlers MustInit"
	var err error

	{
		router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			resultErrJson, _ := result.NewResultErr(error_with_codes.ErrorMethodNotAllowed).GetJson()
			http.Error(w, string(resultErrJson), http.StatusMethodNotAllowed)
		})
	}

	{
		urlGetHandler := url_get_handler.NewUrlGetHandler(usecases.GetUrl())
		initHandler(router, urlGetHandler)
	}

	{
		urlGetListHandler := url_get_list_handler.NewUrlGetListHandler(usecases.GetUrl())
		initHandler(router, urlGetListHandler)
	}

	{
		urlPostHandler := url_post_handler.NewUrlPostHandler(usecases.GetUrl())
		initHandler(router, urlPostHandler)
	}

	return err
}

type handlerI interface {
	GetMethod() handler_type.HandlerMethod
	GetPath() handler_type.HandlerPath
	ExecFunc(ctx context.Context, r *http.Request) ([]byte, error)
}

func initHandler(router *mux.Router, h handlerI) {
	const action = "handler "

	l := logrus.WithFields(logrus.Fields{
		"handler": h.GetPath(),
		"method":  h.GetMethod(),
	})

	l.Info(action, "init")

	router.HandleFunc(h.GetPath().String(), func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		ctxToRun, cancel := context.WithTimeout(r.Context(), defaultExecTimeout)
		defer cancel()

		response, err := h.ExecFunc(ctxToRun, r)
		if err != nil {
			resultJson, err := result.NewResultErr(err).GetJson()
			if err != nil {
				l.WithError(err).
					Error(action, "NewResultErr(err).GetJson()")

				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			w.Write(resultJson)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		}
	}).Methods(h.GetMethod().String())
}
