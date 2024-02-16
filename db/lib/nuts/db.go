package nuts

import (
	"github.com/ImPedro29/rinha-backend-2024/db/interfaces"
	"github.com/nutsdb/nutsdb"
	"go.uber.org/zap"
)

func NewDB() interfaces.DB {
	opts := nutsdb.DefaultOptions
	opts.SyncEnable = false // is faster!
	//opts.RWMode = nutsdb.MMap                       // seems better
	opts.SegmentSize = 8 * nutsdb.MB                 // testing
	opts.CommitBufferSize = 50 * nutsdb.MB           // is not game change
	opts.HintKeyAndRAMIdxCacheSize = 100 * nutsdb.MB // testing

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
