package order

import (
	"theshop/internal/mediator"
	"theshop/internal/messages"
)

type OrderService struct {
	mediator *mediator.Mediator
}

func NewOrderService(m *mediator.Mediator) *OrderService {
	return &OrderService{mediator: m}
}

func (s *OrderService) Handle(cmd messages.PlaceOrder) (messages.OrderResult, error) {
	paymentResult, err := mediator.Send[messages.ProcessPayment, messages.PaymentResult](
		s.mediator,
		messages.ProcessPayment{
			OrderID: cmd.OrderID,
			Amount:  cmd.Amount,
		},
	)
	if err != nil {
		return messages.OrderResult{}, err
	}

	return messages.OrderResult{
		OrderID: cmd.OrderID,
		Success: paymentResult.Success,
	}, nil
}
