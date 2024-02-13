package models

type TransactionRequest struct {
	Value       int64  `json:"valor" validate:"required"`
	Type        string `json:"tipo" validate:"required oneof='c''d'"`
	Description string `json:"descricao" validate:"required"`
}

type TransactionResponse struct {
	Limit   int64 `json:"limite"`
	Balance int64 `json:"saldo"`
}
