package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	//cache.RedisConfig
	//postgres.Config

	GRPCServerPort int `env:"GRPC_SERVER_PORT"`
	RestServerPort int `env:"REST_SERVER_PORT" env-default:"8080"`
}

func New() *Config {
	cfg := Config{}
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &cfg
}
