package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/k1v4/url_shortener/pkg/database/postgres"
)

type Config struct {
	postgres.DBConfig

	GRPCServerPort int `env:"GRPC_SERVER_PORT" envDefault:"50051"`
	RestServerPort int `env:"REST_SERVER_PORT" env-default:"8080"`
}

func New() *Config {
	cfg := Config{}

	// считываем конфиг из .env файла
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &cfg
}
