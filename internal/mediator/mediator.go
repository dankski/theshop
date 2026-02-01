package mediator

type Request interface {
	Name() string
}

type Response interface{}

type RequestHandler interface {
	Handle(req Request) (Response, error)
}

type Mediator interface {
	Send(req Request) (Response, error)
	Register(name string, handler RequestHandler)
}
