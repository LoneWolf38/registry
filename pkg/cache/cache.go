package cache

import (
	"sync"
	"time"

	"github.com/LoneWolf38/registry/pkg/protocol"
	"github.com/LoneWolf38/registry/pkg/registry"
)

// registryCache is the db which will be storing the register in an in-memory database
// the indexing will be done on the timestamp(recieving) and the hostname of the server sending the heartbeat
// Based on the opCode sent, different type of operations will be performed on the insertion on the registry
type registryCache struct {
	mu sync.RWMutex
	v  map[index]registry.Record
}

type index struct {
	rts      time.Time
	hostname string
}

var RCache registryCache

func Default() {
	v := make(map[index]registry.Record)
	RCache = registryCache{
		v: v,
	}
}

// insert is for pushing data in to the cache register and updating the index meanwhile
func (r *registryCache) Insert(data protocol.HeartBeat) {
	r.mu.Lock()
	defer r.mu.Unlock()
}

// delete will remove a record from the cache based on different startegies or based on the opcode and update the index
func (r *registryCache) Delete(data protocol.HeartBeat) {
	r.mu.Lock()
	defer r.mu.Unlock()
}

// updateIndex will update the index of the registry cache
func (r *registryCache) updateIndex() {
}
