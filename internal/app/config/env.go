package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ENV struct{}

func (e ENV) Load() {
	errLoad := godotenv.Load()
	if errLoad != nil {
		panic(fmt.Sprintf("ENV LOAD ERROR: %s", errLoad.Error()))
	}

}

func (e ENV) Get(key string) string {
	return os.Getenv(key)
}
