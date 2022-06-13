package app

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

// Config groups config variables
type Config struct {
	DBUri      string
	DBName     string
	ServerPort string
}

// NewConfig creates a new Config struct instance
func NewConfig() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		return nil, errors.New("no env file found")
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		return nil, errors.New("SERVER_PORT env var is empty")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return nil, errors.New("MONGODB_URI env var is empty")
	}

	database := os.Getenv("MONGO_DATABASE")
	if database == "" {
		return nil, errors.New("MONGO_DATABASE env var is empty")
	}

	return &Config{
		DBUri:      uri,
		DBName:     database,
		ServerPort: serverPort,
	}, nil
}
