package url_repository

import (
	"context"
	"url-shortener/internal/common/types/db_types"
	"url-shortener/internal/common/types/error_with_codes"
	"url-shortener/internal/common/types/url_types"
	"url-shortener/internal/model/url_model"
)

type UrlRepository struct {
	adapter adapterI
	queries struct {
		createUrl    db_types.DbQuery
		getListUrl   db_types.DbQuery
		getUrlByCode db_types.DbQuery
	}
}

type adapterI interface {
	Get(ctx context.Context, dest interface{}, query db_types.DbQuery, params ...interface{}) (ok bool, err error)
	Select(ctx context.Context, dest interface{}, query db_types.DbQuery) (ok bool, err error)
	NamedExec(ctx context.Context, query db_types.DbQuery, arg interface{}) error
}

func NewUrlRepository(adapter adapterI) *UrlRepository {
	repo := new(UrlRepository)
	repo.adapter = adapter

	repo.queries.createUrl = db_types.DbQuery("INSERT INTO url (original_url, code) VALUES (:original_url, :code)")
	repo.queries.getListUrl = db_types.DbQuery("SELECT * FROM url")
	repo.queries.getUrlByCode = db_types.DbQuery("SELECT * FROM url WHERE code = :code")
	return repo
}

func (r *UrlRepository) CreateUrl(ctx context.Context, url *url_model.Url) error {

	err := r.adapter.NamedExec(ctx, r.queries.createUrl, url)
	if err != nil {
		err = error_with_codes.ErrorFailedToCreateUrl
	}

	return err
}

func (r *UrlRepository) GetUrlByCode(ctx context.Context, code url_types.UrlCode) (*url_model.Url, bool, error) {
	dest := new(url_model.Url)

	ok, err := r.adapter.Get(ctx, dest, r.queries.getUrlByCode, code)
	if err != nil {
		err = error_with_codes.ErrorFailedToGetUrlByCode
	}

	return dest, ok, err
}

func (r *UrlRepository) GetUrlList(ctx context.Context) ([]*url_model.Url, bool, error) {
	dest := make([]*url_model.Url, 0)

	ok, err := r.adapter.Select(ctx, &dest, r.queries.getListUrl)
	if err != nil {
		err = error_with_codes.ErrorFailedToGetUrlList
	}

	return dest, ok, err
}
