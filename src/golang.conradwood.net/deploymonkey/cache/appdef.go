package cache

import (
	"sync"

	pb "golang.conradwood.net/apis/deploymonkey"
)

type AppDefCache struct {
	sync.Mutex
	cache map[uint32][]*pb.ApplicationDefinition
}

func NewAppDefCache() *AppDefCache {
	return &AppDefCache{
		cache: make(map[uint32][]*pb.ApplicationDefinition),
	}
}

func (ac *AppDefCache) Reset() {
	ac.Lock()
	defer ac.Unlock()
	ac.cache = make(map[uint32][]*pb.ApplicationDefinition)
}
func (ac *AppDefCache) AddAppGroupVersion(version uint32, appdefs []*pb.ApplicationDefinition) {
	ac.Lock()
	defer ac.Unlock()
	ac.cache[version] = appdefs
}
func (ac *AppDefCache) GetAppGroupVersion(version uint32) []*pb.ApplicationDefinition {
	ac.Lock()
	defer ac.Unlock()
	return ac.cache[version]
}
