package url_get_dto

import "url-shortener/internal/common/types/url_types"

type UrlGetDto struct {
	Code url_types.UrlCode `json:"code"`
}
