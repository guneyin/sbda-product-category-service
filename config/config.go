package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

const defaultRpcPort = 3001

type Config struct {
	RpcPort          int    `env:"RPC_PORT"`
	DiscoverySvcAddr string `env:"DISCOVERY_SVC_ADDR"`
	PgHost           string `env:"PG_HOST"`
	PgPort           string `env:"PG_PORT"`
	PgUser           string `env:"PG_USER"`
	PgPwd            string `env:"PG_PWD"`
	RdsHost          string `env:"RDS_CONN_STR"`
	RdsPort          string `env:"RDS_CONN_STR"`
}

func GetConfig() (*Config, error) {
	config := &Config{RpcPort: defaultRpcPort}

	_ = cleanenv.ReadEnv(config)
	_ = cleanenv.ReadConfig(".env", config)

	return config, nil
}
