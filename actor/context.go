package actor

type Context interface {
	BasePart
	Sender
	Spawner
	Messager
}

type Spawner interface {
	Spawn(props *Props) *PID
	SpawnPrefix(props *Props, prefix string) *PID
	SpawnNamed(props *Props, id string) *PID
}

type Sender interface {
	Sender() *PID
	Send(pid *PID, message any)
	SendByName(name string, message any)
	Call(pid *PID, message any)
}

type BasePart interface {
	Children() []*PID
	Response(response any, err error)
}

type InfoPart interface {
	Parent() *PID
	Self() *PID
	Actor() Actor
	ActorSystem() *System
}

type Messager interface {
	Message() any
}
