package models

type TransactionRequest struct {
	Value       int64  `json:"valor" validate:"required,min=1"`
	Type        string `json:"tipo" validate:"required,oneof='c' 'd'"`
	Description string `json:"descricao" validate:"required,max=10"`
}

type TransactionResponse struct {
	Limit   int64 `json:"limite"`
	Balance int64 `json:"saldo"`
}
