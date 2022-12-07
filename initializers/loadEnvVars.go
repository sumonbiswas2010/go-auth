package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	// err := godotenv.Load()
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}

	// s3Bucket := os.Getenv("S3_BUCKET")
	// secretKey := os.Getenv("SECRET_KEY")
}
