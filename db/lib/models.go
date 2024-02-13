package lib

import (
	"github.com/nutsdb/nutsdb"
)

type db struct {
	instance *nutsdb.DB
}

type client struct {
	id      int64
	balance int64
	limit   int64
}
