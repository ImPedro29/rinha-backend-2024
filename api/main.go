package main

import (
	"log"
	"os"
	"runtime"

	"github.com/ImPedro29/rinha-backend-2024/api/routes"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/prefork"
	"go.uber.org/zap"
)

func main() {
	runtime.GOMAXPROCS(8)

	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	port := "4002"
	if newPort, ok := os.LookupEnv("PORT"); ok {
		port = newPort
	}

	zap.L().Info("start listing on port ", zap.String("port", port))

	server := &fasthttp.Server{
		Handler: routes.InitRoutes,
	}

	preforkServer := prefork.New(server)
	if err := preforkServer.ListenAndServe(":" + port); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
}
