package main

import (
	"exam/user-service/config"
	con "exam/user-service/kafka/consumer"
	"exam/user-service/pkg/logger"
	service2 "exam/user-service/service"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.Environment, "user-service")
	service, err := service2.New(cfg, log)
	if err != nil {
		log.Error("error while accessing services", logger.Error(err))
		return
	}

	//kafka\\
	consumer, err := con.NewKafkaConsumer([]string{"kafka:9092"}, "kafka-user", "", service.UserService)
	if err != nil {
		log.Fatal("error while creating a new kafka consumer", logger.Error(err))
	}
	defer consumer.Close()

	go func() {
		consumer.ConsumeMessages(con.ConsumeHandler)
	}()
	// //kafka\\

	service.Run(log, cfg)
}
