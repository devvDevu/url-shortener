package request_url_get

import "url-shortener/internal/common/types/url_types"

type UrlGetDto struct {
	Code url_types.UrlCode `json:"code" validate:"required,min=1,max=8"`
}
