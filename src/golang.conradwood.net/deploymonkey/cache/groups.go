package cache

import (
	"context"
	"sync"

	pb "golang.conradwood.net/apis/deploymonkey"
)

type GroupCache struct {
	sync.Mutex
	cache_entries []*GroupCacheEntry
}
type GroupCacheEntry struct {
	namespace string
	group     *pb.AppGroup
}

func NewGroupCache() *GroupCache {
	return &GroupCache{}
}
func (gc *GroupCache) Reset() {
	gc.Lock()
	defer gc.Unlock()
	gc.cache_entries = gc.cache_entries[:0]
}
func (gc *GroupCache) AddWithNameSpace(ctx context.Context, namespace string, ag *pb.AppGroup) {
	gc.Lock()
	defer gc.Unlock()
	for _, ce := range gc.cache_entries {
		if ce.namespace == namespace {
			ce.group = ag
			return
		}
	}
	gce := &GroupCacheEntry{namespace: namespace, group: ag}
	gc.cache_entries = append(gc.cache_entries, gce)
}
func (gc *GroupCache) ByNameSpace(ctx context.Context, namespace string) *pb.AppGroup {
	gc.Lock()
	defer gc.Unlock()
	for _, ce := range gc.cache_entries {
		if ce.namespace == namespace {
			return ce.group
		}
	}
	return nil
}
