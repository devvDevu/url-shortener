package request_url_post

import "url-shortener/internal/common/types/url_types"

type UrlPostDto struct {
	OriginalUrl url_types.UrlOriginal `json:"original_url" validate:"required,url"`
}
