package config

var defaultConfig = Config{
	Port:      "8000",
	CertFile:  "",
	KeyFile:   "",
	StaticDir: "",
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
