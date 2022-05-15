package config

var DefaultConfig = Config{
	Fiber: Fiber{
		Address:   "127.0.0.1",
		Port:      "8080",
		KeyFile:   "./private.key",
		CertFile:  "./public.crt",
		StaticDir: "./dist",
	},
}

type Fiber struct {
	Address   string `config:"fiber.address"`
	Port      string `config:"fiber.port"`
	KeyFile   string `config:"fiber.key_file"`
	CertFile  string `config:"fiber.cert_file"`
	StaticDir string `config:"fiber.static_dir"`
}

type Config struct {
	Fiber Fiber
}
