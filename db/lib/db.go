package lib

import (
	"github.com/ImPedro29/rinha-backend-2024/db/interfaces"
	"github.com/nutsdb/nutsdb"
	"go.uber.org/zap"
)

func NewDB() interfaces.DB {
	opts := nutsdb.DefaultOptions
	//opts.SyncEnable = false
	instance, err := nutsdb.Open(
		opts,
		nutsdb.WithDir("data"),
	)

	if err != nil {
		zap.L().Panic("failed to open nutsdb", zap.Error(err))
	}

	return &db{
		instance: instance,
	}
}
