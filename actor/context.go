package actor

type Context interface {
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
	Send(pid *PID, message any)
	Call(pid *PID, message any)
}

type Messager interface {
	Message() any
}
