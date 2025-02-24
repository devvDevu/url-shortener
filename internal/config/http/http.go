package http

type HttpConfig struct {
	Addr string `yaml:"addr" env:"HTTP_ADDR" env-default:"0.0.0.0:8080" env-required:"true" env-upd:""`
}

func (h *HttpConfig) GetAddr() string {
	return h.Addr
}
