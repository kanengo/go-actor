package actor

import (
	"sync/atomic"

	"github.com/kanengo/goutil/pkg/collection"

	"github.com/spaolacci/murmur3"

	cmap "github.com/orcaman/concurrent-map/v2"

	"github.com/kanengo/goutil/pkg/utils"
)

type ProcessManager struct {
	SequenceId   uint64
	ActorSystem  *System
	Address      string
	LocalProcess *collection.Slice[cmap.ConcurrentMap[Process]]

	RemoteHandlers *collection.Slice[AddressResolver]
}

type AddressResolver func(*PID) (Process, bool)

func NewProcessManager(actorSystem *System) *ProcessManager {
	pm := &ProcessManager{
		SequenceId:     0,
		ActorSystem:    actorSystem,
		Address:        localAddress,
		LocalProcess:   collection.NewSlice[cmap.ConcurrentMap[Process]](1024),
		RemoteHandlers: collection.NewSlice[AddressResolver](),
	}

	for i := 0; i < pm.LocalProcess.Length(); i++ {
		if err := pm.LocalProcess.Update(i, cmap.New[Process]()); err != nil {
			panic(err)
		}
	}

	return pm
}

const (
	localAddress = "localProcess"
)

func (pm *ProcessManager) NextId() string {
	nextId := atomic.AddUint64(&pm.SequenceId, 1)
	return uint64ToId(nextId)
}

func (pm *ProcessManager) GetProcessBucket(key string) cmap.ConcurrentMap[Process] {
	hash := murmur3.Sum32(utils.StringToSliceBytesUnsafe(key))
	index := int(hash) % pm.LocalProcess.Length()

	ret, err := pm.LocalProcess.At(index)
	if err != nil {
		panic(err)
	}

	return ret
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
	buf[i] = '@'

	return string(buf[i:])
}

func (pm *ProcessManager) Add(process Process, id string) (*PID, bool) {
	bucket := pm.GetProcessBucket(id)

	return &PID{
		Id:      id,
		Address: pm.Address,
	}, bucket.SetIfAbsent(id, process)
}

func (pm *ProcessManager) Remove(pid *PID) {
	bucket := pm.GetProcessBucket(pid.Id)
	bucket.Pop(pid.Id)
}

func (pm *ProcessManager) Get(pid *PID) (Process, bool) {
	if pid == nil {
		return nil, false
	}

	if !pid.isLocal(pm.Address) {
		for _, resolver := range pm.RemoteHandlers.Source() {
			ref, ok := resolver(pid)
			if ok {
				return ref, true
			}
		}
		//TODO
		return nil, false
	}

	bucket := pm.GetProcessBucket(pid.Id)
	ref, ok := bucket.Get(pid.Id)

	if !ok {
		return nil, false
	}

	return ref, true
}
