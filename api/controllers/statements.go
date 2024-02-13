package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ImPedro29/rinha-backend-2024/api/models"
	"github.com/ImPedro29/rinha-backend-2024/shared/pb"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

func (s Controller) Statements(ctx *fasthttp.RequestCtx, clientId string) {
	id, err := strconv.ParseInt(clientId, 10, 64)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusBadRequest)
		if _, err := ctx.WriteString(`{"message": "client id is not a integer"}`); err != nil {
			zap.L().Error("failed to write response", zap.Error(err))
		}
		return
	}

	res, err := s.db.Statement(ctx, &pb.StatementRequest{
		ClientId: id,
	})
	if err != nil {
		if strings.Contains(err.Error(), "key not found") {
			ctx.Response.SetStatusCode(http.StatusNotFound)
			return
		}

		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		if _, err := ctx.WriteString(`{"message": "failed to create transaction"}`); err != nil {
			zap.L().Error("failed to write response", zap.Error(err))
		}
		return
	}

	var txs []models.StatementTransactionsResponseItem
	for _, tx := range res.LastTransactions {
		createdAt, err := time.Parse(time.RFC3339, tx.CreatedAt)
		if err != nil {
			if _, err := ctx.WriteString(`{"message": "failed to parse date"}`); err != nil {
				zap.L().Error("failed to write response", zap.Error(err))
			}
			return
		}

		txs = append(txs, models.StatementTransactionsResponseItem{
			Value:       tx.Value,
			Description: tx.Description,
			Type:        tx.Type.String()[:1],
			CreatedAt:   createdAt,
		})
	}

	date, err := time.Parse(time.RFC3339, res.Balance.Date)
	if err != nil {
		if _, err := ctx.WriteString(`{"message": "failed to parse date"}`); err != nil {
			zap.L().Error("failed to write response", zap.Error(err))
		}
		return
	}

	resBytes, err := json.Marshal(models.StatementResponse{
		Balance: models.StatementBalanceResponse{
			Total: res.Balance.Total,
			Date:  date,
			Limit: res.Balance.Limit,
		},
		LastTransactions: txs,
	})

	if _, err := ctx.Write(resBytes); err != nil {
		zap.L().Error("failed to write response", zap.Error(err))
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
	}
}
