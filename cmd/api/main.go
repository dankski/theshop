// cmd/api/main.go
package main

import (
	"fmt"
	"theshop/internal/mediator"
	"theshop/internal/messages"
	"theshop/internal/order"
	"theshop/internal/payment"
)

func main() {
	m := mediator.New()

	orderSvc := order.NewOrderService(m)
	paymentSvc := &payment.PaymentService{}

	mediator.Register[messages.PlaceOrder, messages.OrderResult](m, orderSvc)
	mediator.Register[messages.ProcessPayment, messages.PaymentResult](m, paymentSvc)

	result, err := mediator.Send[messages.PlaceOrder, messages.OrderResult](
		m,
		messages.PlaceOrder{
			OrderID: "o123",
			Amount:  99,
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v", result)
}
