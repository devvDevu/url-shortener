package url_model

import "url-shortener/internal/common/types/url_types"

type Url struct {
	Id url_types.UrlId `db:"id"`

	Original url_types.UrlOriginal `db:"original_url"`

	Code url_types.UrlCode `db:"code"`
}
