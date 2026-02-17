package payment

import (
	"theshop/internal/messages"
)

type PaymentService struct{}

func (p *PaymentService) Handle(cmd messages.ProcessPayment) (messages.PaymentResult, error) {
	return messages.PaymentResult{
		OrderID: cmd.OrderID,
		Success: true,
	}, nil
}
