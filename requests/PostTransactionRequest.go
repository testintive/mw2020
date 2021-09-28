package requests

type PostTransactionRequest struct {
	UserEmail         string  `json:"user_email"`
	TransactionAmount float32 `json:"transaction_amount"`
}
