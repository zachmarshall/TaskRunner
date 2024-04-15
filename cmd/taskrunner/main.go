package main

import (
	"fmt"

	"github.com/vatusa/taskrunner/internal/logger"
	"github.com/vatusa/taskrunner/pkg/config"
	"github.com/vatusa/taskrunner/pkg/scheduler"
)

func main() {
	logger.Info("Starting JobScheduler")
	cfg := config.New()

	if cfg.RabbitMQConfig == nil {
		logger.Error("RabbitMQ configuration is missing")
		return
	}

	RabbitMQURI := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		cfg.RabbitMQConfig.Username,
		cfg.RabbitMQConfig.Password,
		cfg.RabbitMQConfig.Host,
		cfg.RabbitMQConfig.Port,
	)

	sch := scheduler.NewScheduler(RabbitMQURI, cfg.RabbitMQConfig.Queue)
	sch.Run()
}
