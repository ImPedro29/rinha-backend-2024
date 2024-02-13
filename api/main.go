package main

import (
	"log"
	"os"

	"github.com/ImPedro29/rinha-backend-2024/api/routes"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	port := "4002"
	if newPort, ok := os.LookupEnv("PORT"); ok {
		port = newPort
	}

	zap.L().Info("start listing on port ", zap.String("port", port))

	if err := fasthttp.ListenAndServe(":"+port, routes.InitRoutes); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
}