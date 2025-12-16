package initial

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	if godotenv.Load() != nil {
		log.Fatal("Failed to load enviroment variables")
	}
}
