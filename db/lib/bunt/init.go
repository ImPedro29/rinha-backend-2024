package bunt

import (
	"fmt"

	"github.com/ImPedro29/rinha-backend-2024/db/constants"
	"github.com/tidwall/buntdb"
)

func (s *db) Init() error {
	return s.instance.Update(func(tx *buntdb.Tx) error {
		for _, c := range constants.Clients {
			balanceKey := fmt.Sprintf(`%d-balance`, c.ID)
			limitKey := fmt.Sprintf(`%d-limit`, c.ID)

			if _, _, err := tx.Set(balanceKey, "0", &buntdb.SetOptions{}); err != nil {
				return err
			}

			if _, _, err := tx.Set(limitKey, fmt.Sprintf("%d", c.Limit), &buntdb.SetOptions{}); err != nil {
				return err
			}

			if err := tx.CreateIndex(fmt.Sprintf("%d-txs", c.ID), fmt.Sprintf("%d:*", c.ID), buntdb.IndexString); err != nil {
				return err
			}
		}

		return nil
	})
}
