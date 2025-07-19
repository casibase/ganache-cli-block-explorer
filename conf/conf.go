package conf

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type ContractConfig struct {
	Path string `yaml:"path"`
	Name string `yaml:"name"`
}

type Config struct {
	ServerAddr  string           `yaml:"server_addr"`
	NetworkHost string           `yaml:"network_host"`
	Contracts   []ContractConfig `yaml:"contracts"`
}

func LoadConfig(path string) Config {
	var config Config
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("Failed to decode yaml: %v", err)
	}
	return config
}
