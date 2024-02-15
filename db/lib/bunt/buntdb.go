package bunt

import (
	"github.com/ImPedro29/rinha-backend-2024/db/interfaces"
	"github.com/tidwall/buntdb"
	"go.uber.org/zap"
)

func NewDB() interfaces.DB {
	conn, err := buntdb.Open(":memory:")
	if err != nil {
		zap.L().Panic("failed to open buntdb", zap.Error(err))
	}

	return &db{
		instance: conn,
	}
}
