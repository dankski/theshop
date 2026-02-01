package mediator

import "fmt"

type Orchestrator struct {
	handlers map[string]RequestHandler
}

func New() *Orchestrator {
	return &Orchestrator{
		handlers: make(map[string]RequestHandler),
	}
}

func (o *Orchestrator) Register(name string, h RequestHandler) {
	o.handlers[name] = h
}

func (o *Orchestrator) Send(req Request) (Response, error) {
	h, ok := o.handlers[req.Name()]
	if !ok {
		return nil, fmt.Errorf("no handler for %s", req.Name())
	}

	return h.Handle(req)
}
