package request_url_post

import "url-shortener/internal/common/types/url_types"

// @Description DTO для создания короткой ссылки
type UrlPostDto struct {
	// Оригинальный URL для сокращения
	// Required: true
	// Example: https://example.com/path?param=value
	// Format: url
	OriginalUrl url_types.UrlOriginal `json:"original_url" validate:"required,url"`
}
