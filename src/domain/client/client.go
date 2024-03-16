package client

import "time"

type ExtractClientInfo struct {
	Balance int       `json:"total"`
	Limit   int       `json:"limite"`
	Date    time.Time `json:"data_extrato"`
}
type ExtractTransaction struct {
	Value       int       `json:"valor"`
	Type        string    `json:"tipo"`
	Description string    `json:"descricao"`
	CreatedAt   time.Time `json:"realizada_em"`
}

type Extract struct {
	Info         ExtractClientInfo    `json:"saldo"`
	Transactions []ExtractTransaction `json:"ultimas_transacoes"`
}

type TransactionReq struct {
	Value       int    `json:"valor" binding:"required"`
	Type        string `json:"tipo" binding:"required"`
	Description string `json:"descricao" binding:"required"`
}

type TransactionRes struct {
	Balance int `json:"saldo"`
	Limit   int `json:"limite"`
}

func (tq *TransactionReq) IsValid() bool {
	return tq.Value > 0 && (len(tq.Description) <= 10 && len(tq.Description) > 0) && (tq.Type == "d" || tq.Type == "c")
}
