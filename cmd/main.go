package main

import (
	"fmt"
	"task_go_api_gateway/api"
	"task_go_api_gateway/api/handlers"
	"task_go_api_gateway/config"
	"task_go_api_gateway/grpc/client"
	"task_go_api_gateway/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	fmt.Printf("config: %+v\n", cfg)

	grpcSvcs, err := client.NewGrpcClients(cfg)
	if err != nil {
		panic(err)
	}

	var loggerLevel = new(string)

	*loggerLevel = logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		*loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.NewLogger("api_gateway", *loggerLevel)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			return
		}
	}()

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	h := handlers.NewHandler(cfg, log, grpcSvcs)

	api.SetUpAPI(r, h, cfg)

	fmt.Println("Start api gateway....")

	if err := r.Run(cfg.ServicePort); err != nil {
		return
	}
}
