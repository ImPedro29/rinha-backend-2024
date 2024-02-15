package nuts

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ImPedro29/rinha-backend-2024/db/constants"
	"github.com/ImPedro29/rinha-backend-2024/shared/pb"
	"github.com/golang/protobuf/proto"
	"github.com/nutsdb/nutsdb"
	"go.uber.org/zap"
)

func (s *db) GetStatement(request *pb.StatementRequest) (*pb.StatementResponse, error) {
	response := &pb.StatementResponse{
		Balance: &pb.Balance{
			Date: time.Now().Format(time.RFC3339),
		},
	}

	if err := s.instance.View(func(tx *nutsdb.Tx) error {
		balanceKey := []byte(fmt.Sprintf(`%d-balance`, request.ClientId))
		limitKey := []byte(fmt.Sprintf(`%d-limit`, request.ClientId))
		transactionsKey := []byte(fmt.Sprintf(`%d-transactions`, request.ClientId))

		result, err := tx.MGet(constants.ClientData, balanceKey, limitKey)
		if err != nil {
			return err
		}

		size, _ := tx.LSize(constants.Transactions, transactionsKey)
		if size > 0 {
			end := size
			if end > 9 {
				end = 9
			}

			txs, err := tx.LRange(constants.Transactions, transactionsKey, 0, end)
			if err != nil {
				zap.L().Error("failed to get txs", zap.Error(err))
			}

			for _, gotTx := range txs {
				var decodedTx pb.TransactionRequest
				if err := proto.Unmarshal(gotTx, &decodedTx); err != nil {
					return err
				}
				response.LastTransactions = append(response.LastTransactions, &pb.LastTransactions{
					Value:       decodedTx.Amount,
					Type:        decodedTx.Type,
					Description: decodedTx.Description,
					CreatedAt:   decodedTx.CreatedAt,
				})
			}
		}

		balance, err := strconv.ParseInt(string(result[0]), 10, 64)
		if err != nil {
			return err
		}

		limit, err := strconv.ParseInt(string(result[1]), 10, 64)
		if err != nil {
			return err
		}

		response.Balance.Total = balance
		response.Balance.Date = time.Now().Format(time.RFC3339)
		response.Balance.Limit = limit

		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
