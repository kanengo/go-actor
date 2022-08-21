package actor

func (pid *PID) isLocal(address string) bool {
	if pid.Address != localAddress && pid.Address != address {
		return false
	}

	return true
}
