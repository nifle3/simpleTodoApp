package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Db         `yaml:"db"`
	HttpServer `yaml:"http_server"`
}

type Db struct {
	Port string `yaml:"port"`
}

type HttpServer struct {
	Port        string        `yaml:"port"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
	JwtKey      string        `yaml:"jwt_key"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	log.Print(configPath)

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("CONFIG_PATH file is not exist")
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal("Can't open file")
	}

	var config Config
	if err = yaml.Unmarshal(file, &config); err != nil {
		log.Fatal("Can't read from file")
	}

	return &config
}
