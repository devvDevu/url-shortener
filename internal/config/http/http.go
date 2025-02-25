package http

type HttpConfig struct {
	Addr     string `yaml:"addr" env:"HTTP_ADDR" env-default:"0.0.0.0:8080" env-required:"true"`
	UseHttps bool   `yaml:"use_https" env:"HTTP_USE_HTTPS" env-default:"false"`
}

func (h *HttpConfig) GetAddr() string {
	return h.Addr
}

func (h *HttpConfig) GetUseHttps() bool {
	return h.UseHttps
}
