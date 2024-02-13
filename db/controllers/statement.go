package controllers

import (
	"context"

	"github.com/ImPedro29/rinha-backend-2024/shared/pb"
)

func (s Controller) Statement(_ context.Context, request *pb.StatementRequest) (*pb.StatementResponse, error) {
	return s.DB.GetStatement(request)
}
