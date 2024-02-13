package models

import (
	"time"
)

type StatementResponse struct {
	Balance          StatementBalanceResponse            `json:"saldo"`
	LastTransactions []StatementTransactionsResponseItem `json:"ultimas_transacoes"`
}

type StatementTransactionsResponseItem struct {
	Value       int64     `json:"valor"`
	Type        string    `json:"tipo"`
	Description string    `json:"descricao"`
	CreatedAt   time.Time `json:"realizada_em"`
}

type StatementBalanceResponse struct {
	Total int64     `json:"total"`
	Date  time.Time `json:"data_extrato"`
	Limit int64     `json:"limite"`
}
