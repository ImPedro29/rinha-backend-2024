package bunt

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/ImPedro29/rinha-backend-2024/shared/pb"
	"github.com/tidwall/buntdb"
	"go.uber.org/zap"
)

func (s *db) GetStatement(request *pb.StatementRequest) (*pb.StatementResponse, error) {
	res := &pb.StatementResponse{
		Balance: &pb.Balance{
			Date: time.Now().Format(time.RFC3339),
		},
	}

	if err := s.instance.View(func(tx *buntdb.Tx) error {
		balanceKey := fmt.Sprintf(`%d-balance`, request.ClientId)
		limitKey := fmt.Sprintf(`%d-limit`, request.ClientId)

		balanceStr, err := tx.Get(balanceKey)
		if err != nil {
			return err
		}
		balance, err := strconv.ParseInt(balanceStr, 10, 64)
		if err != nil {
			return err
		}

		limitStr, err := tx.Get(limitKey)
		if err != nil {
			return err
		}
		limit, err := strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			return err
		}
		res.Balance.Total = balance
		res.Balance.Limit = limit

		i := 0
		err = tx.Ascend(fmt.Sprintf(`%d-txs`, request.ClientId), func(key, value string) bool {
			if i > 9 {
				return false
			}

			var txResult pb.TransactionRequest
			if err := json.Unmarshal([]byte(value), &txResult); err != nil {
				zap.L().Error("failed to unmarshal tx", zap.Error(err))
				return false
			}

			res.LastTransactions = append(res.LastTransactions, &pb.LastTransactions{
				Value:       txResult.Amount,
				Type:        txResult.Type,
				Description: txResult.Description,
				CreatedAt:   txResult.CreatedAt,
			})

			i++
			return true
		})
		if err != nil && !errors.Is(err, buntdb.ErrNotFound) {
			return err
		}

		return nil
	}); err != nil {
		//zap.L().Error("failed to get statement", zap.Error(err))
		return nil, err
	}

	return res, nil
}
