package config

type DatabaseConfig struct {
	Driver                  string
	Url                     string
	ConnMaxLifetimeInMinute uint
	MaxOpenConns            uint
	MaxIdleConns            uint
}

type HttpServerConfig struct {
	Port uint
}
