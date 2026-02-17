package mediator

import (
	"fmt"
	"reflect"
)

type Request[R any] interface {
	isRequest()
}

type Handler[T any, R any] interface {
	Handle(T) (R, error)
}

type Mediator struct {
	handlers map[reflect.Type]any
}

func New() *Mediator {
	return &Mediator{
		handlers: make(map[reflect.Type]any),
	}
}

func Register[T any, R any](m *Mediator, h Handler[T, R]) {
	reqType := reflect.TypeOf((*T)(nil)).Elem()
	m.handlers[reqType] = h
}

func Send[T any, R any](m *Mediator, req T) (R, error) {
	var zero R

	reqType := reflect.TypeOf(req)

	handlerRaw, ok := m.handlers[reqType]
	if !ok {
		return zero, fmt.Errorf("no handler registered for %v", reqType)
	}

	handler, ok := handlerRaw.(Handler[T, R])
	if !ok {
		return zero, fmt.Errorf("handler type mismatch for %v", reqType)
	}

	return handler.Handle(req)
}
