package response_url_get_list

import (
	"url-shortener/internal/common/types/url_types"
	"url-shortener/internal/value_object/url_value_object"
)

type UrlGetListDto struct {
	Urls []Url `json:"urls"`
}

type Url struct {
	Id       url_types.UrlId       `json:"id"`
	Original url_types.UrlOriginal `json:"original_url"`
	Code     url_types.UrlCode     `json:"code"`
}

func NewUrlGetListDto(urlMap map[url_types.UrlId]*url_value_object.Url) *UrlGetListDto {
	urls := make([]Url, len(urlMap))
	i := 0
	for id, url := range urlMap {
		urls[i] = Url{
			Id:       id,
			Original: url.Original,
			Code:     url.Code,
		}
		i++
	}
	return &UrlGetListDto{
		Urls: urls,
	}
}
