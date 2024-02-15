package nuts

import (
	"fmt"
	"strconv"

	"github.com/ImPedro29/rinha-backend-2024/db/constants"
	"github.com/ImPedro29/rinha-backend-2024/shared/common"
	"github.com/ImPedro29/rinha-backend-2024/shared/pb"
	"github.com/golang/protobuf/proto"
	"github.com/nutsdb/nutsdb"
	"go.uber.org/zap"
)

func (s *db) CreateTransaction(request *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	var res pb.TransactionResponse

	if err := s.instance.Update(func(tx *nutsdb.Tx) error {
		// validation limit for debit transactions
		balanceKey := []byte(fmt.Sprintf(`%d-balance`, request.ClientId))
		limitKey := []byte(fmt.Sprintf(`%d-limit`, request.ClientId))
		transactionKey := []byte(fmt.Sprintf(`%d-transactions`, request.ClientId))

		data, err := tx.MGet(constants.ClientData, balanceKey, limitKey)
		if err != nil {
			return err
		}

		balance, err := strconv.ParseInt(string(data[0]), 10, 64)
		if err != nil {
			return err
		}

		limit, err := strconv.ParseInt(string(data[1]), 10, 64)
		if err != nil {
			return err
		}

		amount := request.Amount
		if request.Type == pb.TransactionType_debit {
			amount = -request.Amount
		}

		balance = balance + amount
		if balance < -limit {
			return common.ErrInsufficientBalance
		}

		if err := tx.IncrBy(constants.ClientData, balanceKey, amount); err != nil {
			if err := tx.Rollback(); err != nil {
				zap.L().Error("failed to rollback", zap.Error(err))
			}
			return err
		}

		value, err := proto.Marshal(request)
		if err != nil {
			return err
		}

		if err := tx.LPush(constants.Transactions, transactionKey, value); err != nil {
			if err := tx.Rollback(); err != nil {
				zap.L().Error("failed to rollback", zap.Error(err))
			}
			return err
		}

		res.Balance = balance
		res.Limit = limit

		return nil
	}); err != nil {
		return nil, err
	}

	return &res, nil
}
