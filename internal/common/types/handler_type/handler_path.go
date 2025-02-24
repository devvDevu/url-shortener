package handler_type

type HandlerPath string

func (h HandlerPath) String() string {
	return string(h)
}
