package config

import (
	"encoding/json"
	"os"
)

// LoadConfig 加载配置
func LoadConfig(file string) *Config {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil
	}
	c := new(Config)
	err = json.Unmarshal(bytes, c)
	if err != nil {
		return nil
	}
	return c
}
