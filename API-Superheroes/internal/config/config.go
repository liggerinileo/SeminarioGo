package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// DbConfig ...
type DbConfig struct {
	Driver string `yaml:"driver"`
	Type   string `yaml:"type"`
	Conn   string `yaml:"conn"`
}

// Config ...
type Config struct {
	DB      DbConfig `yaml:"db"`
	Version string   `yaml:"version"`
}

// LoadConfig ...
func LoadConfig(filename string) (*Config, error) {
	// := Creates a variable and infers the type
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config = &Config{}
	yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
