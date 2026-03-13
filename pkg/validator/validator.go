package dotenv_validator

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ValidateEnv(required []string, filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatal("Error while loading .env file")
	}

	missing := []string{}

	for _, key := range required {
		if os.Getenv(key) == "" {
			missing = append(missing, key)
		}
	}

	if len(missing) != 0 {
		log.Fatal("Missing .env vars: ", missing)
	}
}
