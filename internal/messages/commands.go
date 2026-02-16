package messages

type PlaceOrder struct {
	OrderID string
	Amount  float64
}

func (PlaceOrder) Name() string {
	return "PlaceOrder"
}

type ProcessPayment struct {
	OrderID string
	Amount  float64
}

func (ProcessPayment) Name() string { return "ProcessPayment" }
