package main

import (
	"net"
	"os"

	"github.com/ImPedro29/rinha-backend-2024/db/controllers"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	port := "4001"
	if newPort, ok := os.LookupEnv("PORT"); ok {
		port = newPort
	}

	zap.L().Info("start listing on port ", zap.String("port", port))

	listener, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		zap.L().Panic("failed to start server", zap.Error(err))
	}

	controllers.InitControllers(listener)
}
