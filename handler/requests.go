package handler

import "fmt"

func errTransaction(a, t string) error {
	return fmt.Errorf("value %s (type: %s) is required", a, t)
}

func isValidTransactionType(transactionType string) bool {
	return transactionType == "deposit" || transactionType == "withdrawl"
}

type TransactionRequest struct {
	ID              int     `json:"id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

func (t *TransactionRequest) Validate() error {
	if t.Amount == 0 {
		return errTransaction("amount", "float64")
	}
	if t.TransactionType == "" {
		return errTransaction("transaction_type", "string")
	}
	if !isValidTransactionType(t.TransactionType) {
		return fmt.Errorf("value %s must be 'deposit' or 'withdrawl'", t.TransactionType)
	}
	return nil
}
