package config

var defaultConfig = Config{
	Port: "1337",
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
