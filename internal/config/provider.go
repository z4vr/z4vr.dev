package config

var defaultConfig = Config{
	Fiber: Fiber{
		Address:   "127.0.0.1",
		Port:      "8080",
		CertFile:  "public.cert",
		KeyFile:   "private.key",
		StaticDir: "./dist",
	},
}

type Provider interface {
	Load() error
	Instance() *Config
}

type baseProvider struct {
	instance *Config
}

func newBaseProvider() *baseProvider {
	return &baseProvider{
		instance: &defaultConfig,
	}
}

func (p *baseProvider) Instance() *Config {
	return p.instance
}
