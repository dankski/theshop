package payment

import (
	"theshop/internal/mediator"
	"theshop/internal/messages"
)

type Service struct{}

func New() *Service { return &Service{} }

func (s *Service) Handle(req mediator.Request) (mediator.Response, error) {
	cmd := req.(messages.ProcessPayment)

	return messages.PaymentResult{
		OrderID: cmd.OrderID,
		Success: true,
	}, nil
}
