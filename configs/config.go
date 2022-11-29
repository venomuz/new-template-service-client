package configs

import (
	"github.com/spf13/cast"
	"os"
)

type Config struct {
	Environment     string // develop, staging, production
	UserServiceHost string
	UserServicePort int
	CtxTimeout      int
	LogLevel        string
	GrpcServerPort  int
}

func Load() Config {
	c := Config{}
	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "127.0.0.1"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 9000))
	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))
	c.GrpcServerPort = cast.ToInt(getOrReturnDefault("GRPC_SERVER_PORT", 9000))
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
