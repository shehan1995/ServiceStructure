package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

var Config AppConfig

type Service struct {
	ServiceName   string `envconfig:"service_name"`
	Host          string `envconfig:"host"`
	Port          int    `envconfig:"port"`
	ServerTimeout int    `envconfig:"server_timeout"`
	LogLevel      string `envconfig:"log_level"`
}

// AppConfig Config is the main config
type AppConfig struct {
	Service Service `envconfig:"service"`
}

// Load loads the config
func Load() (*AppConfig, error) {
	return LoadWithPrefix("")
}

// LoadWithPrefix loads the config with prefix
func LoadWithPrefix(prefix string) (*AppConfig, error) {
	var config AppConfig
	err := envconfig.Process(prefix, &config)
	return &config, errors.Wrap(err, "process configs")
}

func init() {
	//for local setup
	//os.Setenv("log_level", "debug")
	//defer os.Unsetenv("log_level")
	//
	//os.Setenv("service_name", "qc-promotional-service")
	//defer os.Unsetenv("service_name")
	//
	//os.Setenv("host", "localhost")
	//defer os.Unsetenv("host")
	//
	//os.Setenv("server_timeout", "60")
	//defer os.Unsetenv("server_timeout")
	//
	//os.Setenv("port", "8080")
	//defer os.Unsetenv("port")

	config, err := Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}
	Config = *config
}
