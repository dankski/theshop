package messages

type OrderResult struct {
	OrderID string
	Success bool
}

type PaymentResult struct {
	OrderID string
	Success bool
}
