// cmd/api/main.go
package main

import (
	"net/http"
	"theshop/internal/mediator"
	"theshop/internal/messages"
	"theshop/internal/order"
	"theshop/internal/payment"
	httpTransport "theshop/internal/transport/http"
)

func main() {
	m := mediator.New()

	orderSvc := order.New(m)
	paymentSvc := payment.New()

	// register handlers
	m.Register(messages.PlaceOrder{}.Name(), orderSvc)
	m.Register(messages.ProcessPayment{}.Name(), paymentSvc)

	// REST layer
	h := httpTransport.New(m)

	http.HandleFunc("/orders", h.PlaceOrder)
	http.ListenAndServe(":8080", nil)
}
