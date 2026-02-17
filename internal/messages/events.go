package messages

type OrderResult struct {
	OrderID string `json:"order_id"`
	Success bool   `json:"success"`
}

type PaymentResult struct {
	OrderID string `json:"order_id"`
	Success bool   `json:"success"`
}
