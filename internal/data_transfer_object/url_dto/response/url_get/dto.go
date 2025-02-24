package url_get_dto

import "url-shortener/internal/common/types/url_types"

type UrlGetDto struct {
	OriginalUrl url_types.UrlOriginal `json:"original_url"`
}
