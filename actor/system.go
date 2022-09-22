package actor

import (
	"net"
	"strconv"

	"github.com/lithammer/shortuuid/v4"
)

type System struct {
	ProcessManager *ProcessManager
	Root           Context
	Timeout        Context
	ID             string
	Config         *Config
}

func (as *System) GetHostPort() (host string, port int, err error) {
	addr := as.ProcessManager.Address
	if h, p, e := net.SplitHostPort(addr); e != nil {
		if addr != localAddress {
			err = e
		}

		host = localAddress
		port = -1
	} else {
		host = h
		port, err = strconv.Atoi(p)
	}

	return
}

func NewActorSystem(opts ...ConfigOption) *System {
	config := Configure(opts...)
	return NewActorSystemWithConfig(config)
}

func NewActorSystemWithConfig(config *Config) *System {
	system := &System{
		ProcessManager: nil,
		Root:           nil,
		Timeout:        nil,
		ID:             shortuuid.New(),
		Config:         config,
	}

	system.ProcessManager = NewProcessManager(system)

	return system
}
