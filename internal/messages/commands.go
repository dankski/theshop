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

func (p ProcessPayment) Name() string {
	return "ProcessPayment"
}

type PaymentProcessed struct {
	OrderID string
	Success bool
}

func (p PaymentProcessed) Name() string {
	return "PaymentProcessed"
}

type CreateOrder struct {
	OrderID string
	Amount  float64
}

func (c CreateOrder) Name() string {
	return "CreateOrder"
}
