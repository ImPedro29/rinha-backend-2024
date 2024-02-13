package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/ImPedro29/rinha-backend-2024/api/models"
	"github.com/ImPedro29/rinha-backend-2024/db/lib"
	"github.com/ImPedro29/rinha-backend-2024/shared/pb"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

func (s Controller) CreateTransaction(ctx *fasthttp.RequestCtx, clientId string) {
	id, err := strconv.ParseInt(clientId, 10, 64)
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusBadRequest)
		if _, err := ctx.WriteString(`{"message": "client id is not a integer"}`); err != nil {
			zap.L().Error("failed to write response", zap.Error(err))
		}
		return
	}

	var request models.TransactionRequest
	if err := json.Unmarshal(ctx.Request.Body(), &request); err != nil {
		ctx.Response.SetStatusCode(http.StatusBadRequest)
		if _, err := ctx.WriteString(`{"message": "invalid data sent"}`); err != nil {
			zap.L().Error("failed to write response", zap.Error(err))
		}
		return
	}

	rType := pb.TransactionType_credit
	if request.Type == "d" {
		rType = pb.TransactionType_debit
	}

	res, err := s.db.CreateTransaction(ctx, &pb.TransactionRequest{
		ClientId:    id,
		Amount:      request.Value,
		Type:        rType,
		Description: request.Description,
	})
	if err != nil {
		if strings.Contains(err.Error(), lib.ErrInsufficientBalance.Error()) {
			ctx.Response.SetStatusCode(http.StatusUnprocessableEntity)
			if _, err := ctx.WriteString(`{"message": "insufficient balance"}`); err != nil {
				zap.L().Error("failed to write response", zap.Error(err))
			}
			return
		}

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

	resBytes, err := json.Marshal(models.TransactionResponse{
		Limit:   res.Limit,
		Balance: res.Balance,
	})

	if _, err := ctx.Write(resBytes); err != nil {
		zap.L().Error("failed to write response", zap.Error(err))
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
	}
}
