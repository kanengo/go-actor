package actor

import (
	"sync/atomic"

	cmap "github.com/orcaman/concurrent-map/v2"
)

type ProcessManager struct {
	SequenceId uint64
	*ActorSystem
	Adress       string
	LocalProcess cmap.ConcurrentMap[Process]
}

const (
	localAddress = "localProcess"
)

func (pm *ProcessManager) NextId() string {
	nextId := atomic.AddUint64(&pm.SequenceId, 1)
	return uint64ToId(nextId)
}

const (
	digits = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~+"
)

func uint64ToId(u uint64) string {
	var buf [13]byte
	i := 13
	// base is power of 2: use shifts and masks instead of / and %
	for u >= 64 {
		i--
		buf[i] = digits[uintptr(u)&0x3f]
		u >>= 6
	}
	// u < base
	i--
	buf[i] = digits[uintptr(u)]
	i--
	buf[i] = '$'

	return string(buf[i:])
}
