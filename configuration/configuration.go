package configuration

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"sync"
)

type Configuration struct {
	UrlsFile string `envconfig:"urls_file" required:"true"`
	HTTPport string `envconfig:"http_port" required:"true"`
}

var (
	cfg  *Configuration
	once sync.Once
)

func LoadConfig() {
	once.Do(func() {
		if err := godotenv.Load(".env"); err != nil {
			panic(fmt.Errorf("failed to open .env file: %w", err))
		}

		var config Configuration
		if err := envconfig.Process("", &config); err != nil {
			panic(fmt.Errorf("failed to process .env file: %w", err))
		}
		cfg = &config
	})
}

func GetConfiguration() *Configuration {
	if cfg == nil {
		LoadConfig()
	}
	return cfg
}
