package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_USER", "user")
	os.Setenv("POSTGRES_PASSWORD", "password")
	os.Setenv("POSTGRES_DB", "dbname")
	os.Setenv("GRPC_SERVER_PORT", "50051")
	os.Setenv("REST_SERVER_PORT", "8080")
	os.Setenv("PG_POOL_MAX", "10")

	cfg := New()

	assert.NotNil(t, cfg)

	assert.Equal(t, "localhost", cfg.DBConfig.Host)
	assert.Equal(t, "5432", cfg.DBConfig.Port)
	assert.Equal(t, "user", cfg.DBConfig.UserName)
	assert.Equal(t, "password", cfg.DBConfig.Password)
	assert.Equal(t, "dbname", cfg.DBConfig.DbName)
	assert.Equal(t, 50051, cfg.GRPCServerPort)
	assert.Equal(t, 8080, cfg.RestServerPort)
}

func TestNewConfig_MissingRequiredEnv(t *testing.T) {
	os.Clearenv()

	cfg := New()

	assert.Nil(t, cfg)
}
