package config

import "time"

type Config struct {
	RequestTimeout time.Duration
	GlobalTimeout  time.Duration
	ResultsFile    string
	MaxRetries     int
	BufferSize     int
}

// Редактируемые параметры для запуска
func Load() *Config {
	return &Config{
		RequestTimeout: 2 * time.Second,
		GlobalTimeout:  5 * time.Second,
		ResultsFile:    "results.txt",
		MaxRetries:     3,
		BufferSize:     100,
	}
}
