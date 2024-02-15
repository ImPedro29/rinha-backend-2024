package controllers

import (
	"net"

	"github.com/ImPedro29/rinha-backend-2024/db/interfaces"
	"github.com/ImPedro29/rinha-backend-2024/db/lib/nuts"
	"github.com/ImPedro29/rinha-backend-2024/shared/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Controller struct {
	pb.BankServiceServer
	DB interfaces.DB
}

func InitControllers(listener net.Listener) {
	server := grpc.NewServer()

	db := nuts.NewDB()
	if err := db.Init(); err != nil {
		zap.L().Panic("failed to initialize db", zap.Error(err))
	}

	pb.RegisterBankServiceServer(server, &Controller{
		DB: db,
	})

	if err := server.Serve(listener); err != nil {
		zap.L().Panic("failed to start grpc server", zap.Error(err))
	}
}
