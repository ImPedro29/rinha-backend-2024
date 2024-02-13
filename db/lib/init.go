package lib

import (
	"errors"
	"fmt"

	"github.com/ImPedro29/rinha-backend-2024/db/constants"
	"github.com/nutsdb/nutsdb"
)

func (s *db) Init() error {
	if err := s.instance.Update(func(tx *nutsdb.Tx) error {
		if err := tx.NewListBucket(constants.Transactions); err != nil && !errors.Is(err, nutsdb.ErrBucketAlreadyExist) {
			return err
		}

		if err := tx.NewKVBucket(constants.ClientData); err != nil && !errors.Is(err, nutsdb.ErrBucketAlreadyExist) {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return s.instance.Update(func(tx *nutsdb.Tx) error {
		for _, client := range clients {
			if err := tx.PutIfNotExists(
				constants.ClientData,
				[]byte(fmt.Sprintf(`%d-balance`, client.id)),
				[]byte(`0`),
				0); err != nil {
				return err
			}
			if err := tx.PutIfNotExists(
				constants.ClientData,
				[]byte(fmt.Sprintf(`%d-limit`, client.id)),
				[]byte(fmt.Sprintf("%d", client.limit)), 0); err != nil {
				return err
			}
		}

		return nil
	})

}
