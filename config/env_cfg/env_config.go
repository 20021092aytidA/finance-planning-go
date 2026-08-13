package env_cfg

import (
	"os"

	"github.com/joho/godotenv"
)

type ENVVar struct {
}

var ENVGlobal ENVVar

func Load() error {
	return godotenv.Load()
}

func Get(key string) string {
	return os.Getenv(key)
}
