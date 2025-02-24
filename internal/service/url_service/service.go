package url_service

import (
	"context"
	"url-shortener/internal/common/types/url_types"
	"url-shortener/internal/model/url_model"
)

// future: add cache
type UrlService struct {
	repository repositoryI
}

type repositoryI interface {
	CreateUrl(ctx context.Context, url *url_model.Url) error
	GetUrlList(ctx context.Context) ([]*url_model.Url, bool, error)
	GetUrlByCode(ctx context.Context, code url_types.UrlCode) (*url_model.Url, bool, error)
}

func NewUrlService(repository repositoryI) *UrlService {
	return &UrlService{repository: repository}
}

func (s *UrlService) CreateUrl(ctx context.Context, url *url_model.Url) error {
	return s.repository.CreateUrl(ctx, url)
}

func (s *UrlService) GetUrlList(ctx context.Context) ([]*url_model.Url, bool, error) {
	return s.repository.GetUrlList(ctx)
}

func (s *UrlService) GetUrlByCode(ctx context.Context, code url_types.UrlCode) (*url_model.Url, bool, error) {
	return s.repository.GetUrlByCode(ctx, code)
}
