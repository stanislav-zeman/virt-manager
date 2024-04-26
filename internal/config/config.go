package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Title      string `yaml:"title"`
		Height     int    `yaml:"height"`
		Width      int    `yaml:"width"`
		Background struct {
			Colour struct {
				Red   uint8 `yaml:"red"`
				Green uint8 `yaml:"green"`
				Blue  uint8 `yaml:"blue"`
				Alpha uint8 `yaml:"alpha"`
			} `yaml:"colour"`
		} `yaml:"background"`
	} `yaml:"app"`
}

func Parse(path string) *Config {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil
	}

	var conf *Config
	err = yaml.Unmarshal(bytes, &conf)
	if err != nil {
		return nil
	}

	return conf
}
