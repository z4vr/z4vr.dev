package config

var DefaultConfig = Config{
	Address:   "127.0.0.1",
	Port:      "8080",
	StaticDir: "dist",
	IndexFile: "index.html",
	KeyFile:   "",
	CertFile:  "",
}

type Config struct {
	Address string
	Port    string

	StaticDir string
	IndexFile string

	KeyFile  string
	CertFile string
}
