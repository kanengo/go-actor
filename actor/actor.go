package actor

type Producer func() Actor

type Actor interface {
	Receive(ctx Context)
}
