package http

import (
	"encoding/json"
	"net/http"
	"theshop/internal/mediator"
	"theshop/internal/messages"
)

type Handler struct {
	mediator mediator.Mediator
}

func New(m mediator.Mediator) *Handler {
	return &Handler{mediator: m}
}

func (h *Handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var req messages.PlaceOrder
	_ = json.NewDecoder(r.Body).Decode(&req)

	resp, err := h.mediator.Send(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
