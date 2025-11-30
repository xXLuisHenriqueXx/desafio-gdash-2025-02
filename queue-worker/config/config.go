package config

import (
	"os"
	"time"
)

type Config struct {
	RabbitURL   string
	QueueName   string
	NestAPIURL  string
	HTTPTimeout time.Duration
	MaxRetries  int
}

func Load() Config {
	return Config{
		RabbitURL:   getEnv("RABBITMQ_URL", "amqp://guest:guest@rabbitmq:5672/"),
		QueueName:   getEnv("QUEUE_NAME", "weather.data"),
		NestAPIURL:  getEnv("NEST_API_URL", "http://nestjs:3000/api/weather/logs"),
		HTTPTimeout: 5 * time.Second,
		MaxRetries:  3,
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
