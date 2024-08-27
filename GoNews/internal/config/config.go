package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// Config - структура конфигурации
type Config struct {
	URLS        []string `yaml:"rss"`
	Period      int      `yaml:"request_period"`
	StoragePath string   `yaml:"storage_path"`
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
