package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	HttpConfig           HttpConfig
	AuthenticationConfig AuthenticationConfig
}

type HttpConfig struct {
	Port         string
	IsProduction bool
	Timeout      int
}
type AuthenticationConfig struct {
	GoogleClientId     string
	GoogleClientSecret string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, falling back to environment variables")
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	// Read environment variables
	port := getEnvVariable("PORT", ":8080")
	isProduction := getEnvAsBool("ISPRODUCTION", false)
	timeout := getEnvVariableAsInt("TIMEOUT", 20)
	googleClientId := getEnvVariable("GOOGLE_CLIENT_ID", "no client id")
	googleClientSecret := getEnvVariable("GOOGLE_CLIENT_SECRET", "no client secret")

	return &Config{
		HttpConfig: HttpConfig{
			Port:         port,
			IsProduction: isProduction,
			Timeout:      timeout,
		},
		AuthenticationConfig: AuthenticationConfig{
			GoogleClientId:     googleClientId,
			GoogleClientSecret: googleClientSecret,
		},
	}, nil
}

// Helper functions
func getEnvVariable(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvVariableAsInt(name string, defaultVal int) int {
	if valueStr, exists := os.LookupEnv(name); exists {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}
	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	if valStr, exists := os.LookupEnv(name); exists {
		if val, err := strconv.ParseBool(valStr); err == nil {
			return val
		}
	}
	return defaultVal
}
