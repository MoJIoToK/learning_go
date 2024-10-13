package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Config - Структура конфига
type Config struct {
	CensorList []string `yaml:"censor_list"`
	HTTPServer `yaml:"http_server"`
}

// HTTPServer - структура с настройками сервера
type HTTPServer struct {
	Address      string        `yaml:"address"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
}

// MustLoad - инициализирует данные из конфиг файла.
func MustLoad(path string) *Config {
	var cfg Config

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", path)
	}

	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &cfg
}
