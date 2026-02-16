package mediator

import "fmt"

type SimpleMediator struct {
	handlers map[string]RequestHandler
}

func New() *SimpleMediator {
	return &SimpleMediator{
		handlers: make(map[string]RequestHandler),
	}
}

func (m *SimpleMediator) Register(name string, h RequestHandler) {
	m.handlers[name] = h
}

func (m *SimpleMediator) Send(req Request) (Response, error) {
	h, ok := m.handlers[req.Name()]
	if !ok {
		return nil, fmt.Errorf("no handler for %s", req.Name())
	}
	return h.Handle(req)
}
