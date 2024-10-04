package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

type Config struct {
	Application
	HttpServer HTTPServer `yaml:"http_server" env-required:"true"`
	Database   Database   `yaml:"database" env-required:"true"`
}

type Application struct {
	Env string `yaml:"env" env-required:"true"`
}

type HTTPServer struct {
	Address         string        `yaml:"address" env-default:"localhost"`
	PublicPort      int           `yaml:"public_port" env-default:"8080"`
	PrivatePort     int           `yaml:"private_port" env-default:"8081"`
	Timeout         time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout     time.Duration `yaml:"idle_timeout" env-default:"60s"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" env-default:"10s"`
}

type Database struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     int    `yaml:"port" env-default:"5432"`
	Name     string `yaml:"name" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
}

func MustLoad() *Config {
	configPath := getConfigPath()
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	cfg := &Config{}

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		log.Fatalf("Error reading config: %s", err)
	}

	return cfg
}

func getConfigPath() string {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config/local.yaml"
	}
	return configPath
}
