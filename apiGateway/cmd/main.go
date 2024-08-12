package main

import (
	"api/api"
	"api/api/handler"
	"api/configs"
	"api/grpc/client"
	logger "api/pkg"
	"api/rabbitMq"
	"log"
	"log/slog"
)

func main() {
	cfg := configs.Load()

	logg := logger.NewLogger()

	services, err := client.NewGrpcClients(cfg)
	if err != nil {
		logg.Error("Error while initializing grpc clients", err)
		return
	}
	writer, err := rabbitMq.NewRabbitMqProducerInt("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println(err)
		logg.Error("this error is newRabbitMq", err)
		return
	}

	h := handler.New(cfg, services, *logg, writer)

	r := api.New(h)

	logg.Info("Server is running ...", slog.Any("port", cfg.HTTPort))
	if err = r.Run(cfg.HTTPort); err != nil {
		logg.Error("Error while running server", err)
	}
}
