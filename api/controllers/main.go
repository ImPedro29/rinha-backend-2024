package controllers

import (
	"os"

	"github.com/ImPedro29/rinha-backend-2024/shared/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Controller struct {
	db pb.BankServiceClient
}

func NewController() Controller {
	conn, err := grpc.Dial(os.Getenv("DB_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		zap.L().Panic("failed to connect to db", zap.Error(err))
	}

	return Controller{
		pb.NewBankServiceClient(conn),
	}
}
