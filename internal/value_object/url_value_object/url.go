package url_value_object

import "url-shortener/internal/common/types/url_types"

type Url struct {
	Original url_types.UrlOriginal

	Code url_types.UrlCode
}

func NewUrl(original url_types.UrlOriginal) *Url {
	return &Url{
		Original: original,
		Code:     url_types.UrlCode(original),
	}
}
