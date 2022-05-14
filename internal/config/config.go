package config

type Config struct {
	Port     string `config:"port"`
	CertFile string `config:"cert_file"`
	KeyFile  string `config:"key_file"`
}
