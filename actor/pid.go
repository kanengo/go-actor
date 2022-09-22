package actor

func (x *PID) isLocal(address string) bool {
	if x.Address != localAddress && x.Address != address {
		return false
	}

	return true
}
