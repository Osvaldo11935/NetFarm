package extensions

import (
	"github.com/joho/godotenv"
	"os"
)

func GetEnv(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
	}

	return os.Getenv(key)
}
