package order

import (
	"theshop/internal/mediator"
	"theshop/internal/messages"
)

type Service struct {
	mediator mediator.Mediator
}

func New(m mediator.Mediator) *Service {
	return &Service{mediator: m}
}

func (s *Service) Handle(req mediator.Request) (mediator.Response, error) {
	cmd := req.(messages.PlaceOrder)

	// call payment via mediator → cross-cutting orchestration
	resp, err := s.mediator.Send(messages.ProcessPayment{
		OrderID: cmd.OrderID,
		Amount:  cmd.Amount,
	})
	if err != nil {
		return nil, err
	}

	payResult := resp.(messages.PaymentResult)

	return messages.OrderResult{
		OrderID: cmd.OrderID,
		Success: payResult.Success,
	}, nil
}
