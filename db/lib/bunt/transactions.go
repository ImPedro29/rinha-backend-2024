package bunt

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/ImPedro29/rinha-backend-2024/shared/common"
	"github.com/ImPedro29/rinha-backend-2024/shared/pb"
	"github.com/tidwall/buntdb"
)

func (s *db) CreateTransaction(request *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	var res pb.TransactionResponse

	if err := s.instance.Update(func(tx *buntdb.Tx) error {
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

		if request.Type == pb.TransactionType_debit {
			balance = balance - request.Amount
		} else {
			balance = balance + request.Amount
		}

		if balance < -limit {
			return common.ErrInsufficientBalance
		}

		res.Balance = balance
		res.Limit = limit

		if _, _, err = tx.Set(balanceKey, fmt.Sprintf("%d", balance), &buntdb.SetOptions{}); err != nil {
			if err := tx.Rollback(); err != nil {
				return err
			}
			return err
		}

		txKey := fmt.Sprintf(`%d:%d`, request.ClientId, time.Now().Unix())
		txBytes, err := json.Marshal(request)
		if err != nil {
			if err := tx.Rollback(); err != nil {
				return err
			}
			return err
		}

		_, _, err = tx.Set(txKey, string(txBytes), &buntdb.SetOptions{})

		return err
	}); err != nil {
		//zap.L().Error("failed to create transaction", zap.Error(err))
		return nil, err
	}

	return &res, nil
}
