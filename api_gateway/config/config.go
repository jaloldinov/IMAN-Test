package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	// post GRUD service
	SecondServiceHost string
	SecondServicePort int

	// Data service
	FirstServiceHost string
	FirstServicePort int

	LogLevel string
	HttpPort string
}

// Load loads environment vars and inflates Config
func Load() Config {
	// if err := godotenv.Load(); err != nil {
	// 	fmt.Println("No .env file found")
	// }

	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	config.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))

	// Data service
	config.FirstServiceHost = cast.ToString(getOrReturnDefault("SERVICE_HOST", "localhost"))
	config.FirstServicePort = cast.ToInt(getOrReturnDefault("SERVICE_PORT", 9001))

	// post GRUD service
	config.SecondServiceHost = cast.ToString(getOrReturnDefault("SERVICE_HOST", "localhost"))
	config.SecondServicePort = cast.ToInt(getOrReturnDefault("SERVICE_PORT", 9002))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
