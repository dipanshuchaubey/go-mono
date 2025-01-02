package config

import (
	"fmt"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

const (
	ENV_LOCAL = "local"
	ENV_DEV   = "dev"
	ENV_STAGE = "stage"
	ENV_PROD  = "prod"
)

// LoadConfig loads the configuration for the service
// serviceName: the name of the service
// config: the configuration struct to load the configuration into (passed by reference)
func LoadConfig(serviceName string, config any) {
	env := os.Getenv("ENV")
	if env == "" {
		fmt.Println("ENV not set, defaulting to local")
		env = ENV_LOCAL
	}

	switch env {
	case ENV_LOCAL:
		env = ENV_LOCAL
	case ENV_DEV:
		env = ENV_DEV
	case ENV_STAGE:
		env = ENV_STAGE
	default:
		log.Fatalf("invalid environment: %s", env)
	}

	configFile, fileErr := os.ReadFile(fmt.Sprintf("services/%s/config/%s/config.yaml", serviceName, env))
	if fileErr != nil {
		log.Fatalf("failed to read config file: %v", fileErr)
	}

	if yamlErr := yaml.Unmarshal(configFile, config); yamlErr != nil {
		log.Fatalf("failed to unmarshal config file: %v", yamlErr)
	}

	fmt.Printf("%s config loaded successfully...\n", env)
}
