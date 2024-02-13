package interfaces

import (
	"github.com/ImPedro29/rinha-backend-2024/shared/pb"
)

type DB interface {
	Init() error
	CreateTransaction(*pb.TransactionRequest) (*pb.TransactionResponse, error)
	GetStatement(*pb.StatementRequest) (*pb.StatementResponse, error)
}
