package config

import "os"

type Config struct {
	RabbitMQConfig *RabbitMQConfig
}

type RabbitMQConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Queue    string
}

func RabbitMQ() *RabbitMQConfig {
	return &RabbitMQConfig{
		Host:     os.Getenv("RABBITMQ_HOST"),
		Port:     os.Getenv("RABBITMQ_PORT"),
		Username: os.Getenv("RABBITMQ_USERNAME"),
		Password: os.Getenv("RABBITMQ_PASSWORD"),
		Queue:    os.Getenv("RABBITMQ_QUEUE"),
	}
}

func New() *Config {
	return &Config{
		RabbitMQConfig: RabbitMQ(),
	}
}
