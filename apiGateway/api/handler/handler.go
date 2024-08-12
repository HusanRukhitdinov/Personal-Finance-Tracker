package handler

import (
	"api/api/models"
	config "api/configs"
	"api/grpc/client"
	"api/rabbitMq"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg              config.Config
	log              slog.Logger
	services         client.IServiceManager
	rabbitMqProducer *rabbitMq.RabbitMqProducerInt
}

func New(cfg config.Config, services client.IServiceManager, log slog.Logger,rabbitMqProducer *rabbitMq.RabbitMqProducerInt) Handler {
	return Handler{
		cfg:      cfg,
		services: services,
		log:      log,
		rabbitMqProducer: rabbitMqProducer,
	}
}

func handleResponse(c *gin.Context, log slog.Logger, msg string, statusCode int, data interface{}) {

	resp := models.Response{}

	switch code := statusCode; {
	case code < 400:
		resp.Description = "OK"
		log.Info("~~~~> OK", slog.String("msg", msg), slog.Any("status", code))
	case code == 401:
		resp.Description = "Unauthorized"
		log.Error("???? Unauthorized", slog.String("msg", msg), slog.Any("status", code))
	case code < 500:
		resp.Description = "Bad Request"
		log.Error("!!!!! BAD REQUEST", slog.String("msg", msg), slog.Any("status", code))
	default:
		resp.Description = "Internal Server Error"
		log.Error("!!!!! INTERNAL SERVER ERROR", slog.String("msg", msg), slog.Any("status", code), slog.Any("error", data))
	}

	resp.StatusCode = statusCode
	resp.Data = data

	c.JSON(resp.StatusCode, resp)

}
