package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"
`
	Database struct {
		Url string `yaml:"url"`
	} `yaml:"database"`
}

func LoadConfig(path string) *Config {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open config file: %v", err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatalf("Could not parse config file: %v", err)
	}
	log.Println("Loaded config:", cfg)
	return &cfg
}
