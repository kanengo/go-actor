package actor

type Process interface {
	SendUserMessage(pid *PID, message any)
	SendSystemMessage(pid *PID, message any)
	Stop(pid *PID)
}
