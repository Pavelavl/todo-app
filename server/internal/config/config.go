package config

import (
	"os"
	"sync"
	"todo-app/pkg/logging"

	"github.com/golobby/dotenv"
)

type Config struct {
	IsDebug bool `env:"IS_DEBUG"`
	Listen  struct {
		BindIP string `env:"BIND_IP"`
		Port   string `env:"PORT"`
		Type   string `env:"TYPE"`
	}
	Database struct {
		Name       string `env:"DB_NAME"`
		Host       string `env:"DB_HOST"`
		Collection string `env:"DB_COLLECTION"`
		User       string `env:"DB_USER"`
		Password   string `env:"DB_PASSWORD"`
		AuthDB     string `env:"DB_AUTH"`
		Port       string `env:"DB_PORT"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	config := Config{}
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("Read application configuration")
		instance = &config
		file, err := os.Open(".env")
		if err != nil {
			logger.Fatal(err)
		}
		if err = dotenv.NewDecoder(file).Decode(&config); err != nil {
			logger.Fatal(err)
		}
	})
	return instance
}
