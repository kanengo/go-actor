package actor

import (
	"sync"
)

type ActorContext struct {
	system   *System
	children []*PID
	sender   *PID
}

func (ac *ActorContext) Children() []*PID {
	return ac.children
}

func (ac *ActorContext) Response(response any, err error) {
	//TODO implement me
	panic("implement me")
}

func (ac *ActorContext) Sender() *PID {
	return ac.sender
}

func (ac *ActorContext) Send(pid *PID, message any) {
	//TODO implement me
	panic("implement me")
}

var PIDPool = &sync.Pool{
	New: func() any {
		return new(PID)
	},
}

func (ac *ActorContext) SendByName(name string, message any) {
	pid := PIDPool.Get().(*PID)
	pid.Id = name
	defer func() {
		pid.Id = ""
		pid.Address = ""
		PIDPool.Put(pid)
	}()
}

func (ac *ActorContext) Call(pid *PID, message any) {
	//TODO implement me
	panic("implement me")
}

func (ac *ActorContext) Spawn(props *Props) *PID {
	//TODO implement me
	panic("implement me")
}

func (ac *ActorContext) SpawnPrefix(props *Props, prefix string) *PID {
	//TODO implement me
	panic("implement me")
}

func (ac *ActorContext) SpawnNamed(props *Props, id string) *PID {
	//TODO implement me
	panic("implement me")
}

func (ac *ActorContext) Message() any {
	//TODO implement me
	panic("implement me")
}

func NewActorContext(system *System) *ActorContext {
	actorCtx := &ActorContext{
		system: system,
	}

	return actorCtx
}
