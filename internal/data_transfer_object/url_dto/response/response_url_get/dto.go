package response_url_get

import "url-shortener/internal/common/types/url_types"

type UrlGetDto struct {
	OriginalUrl url_types.UrlOriginal `json:"original_url"`
}
