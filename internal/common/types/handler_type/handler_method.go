package handler_type

import "net/http"

type HandlerMethod string

const (
	GET HandlerMethod = http.MethodGet
	POST HandlerMethod = http.MethodPost
	PUT HandlerMethod = http.MethodPut
	DELETE HandlerMethod = http.MethodDelete
)	

func (h HandlerMethod) String() string {
	return string(h)
}
