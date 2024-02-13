package controllers

import (
	"context"
	"time"

	"github.com/ImPedro29/rinha-backend-2024/shared/pb"
)

func (s Controller) CreateTransaction(_ context.Context, request *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	request.CreatedAt = time.Now().Format(time.RFC3339)

	return s.DB.CreateTransaction(request)
}
