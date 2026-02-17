package messages

type PlaceOrder struct {
	OrderID string
	Amount  float64
}

type ProcessPayment struct {
	OrderID string
	Amount  float64
}
