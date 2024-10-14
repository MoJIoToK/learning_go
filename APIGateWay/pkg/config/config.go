package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

// Config - структура конфигурации
type Config struct {
	News       string `yaml:"news_service"`
	Comments   string `yaml:"comments_service"`
	Censor     string `yaml:"censor_service"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address      string        `yaml:"address"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
}

// MustLoad - инициализирует данные из файла конфигурации. Путь к файлу передаётся в из функции main.
// Если не удается раскодировать файл, то приложение завершается с ошибкой.
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
