package gomap_concurrent

import "sync"

type ShardResolver func(key interface{}, sCount uint16) uint

type GConcMap struct {
	shards     []*GMapShard
	Resolver   ShardResolver
	ShardCount uint16
}

type GMapShard struct {
	items map[interface{}]interface{}
	sync.RWMutex
}

func NewMapWithResolver(resolver ShardResolver, shardCount uint16) GConcMap {
	cm := make([]*GMapShard, shardCount)
	for i := 0; i < int(shardCount); i++ {
		cm[i] = &GMapShard{items: make(map[interface{}]interface{})}
	}
	s := GConcMap{
		Resolver:   resolver,
		ShardCount: shardCount,
		shards:     cm,
	}
	return s
}

func (cm GConcMap) Get(key interface{}) (interface{}, bool) {
	shard := cm.GetShard(key)

	shard.RLock()
	value, ok := shard.items[key]
	shard.RUnlock()

	return value, ok
}
func (cm GConcMap) Set(key interface{}, value interface{}) {
	shard := cm.GetShard(key)

	shard.Lock()
	shard.items[key] = value
	shard.Unlock()
}

func (cm GConcMap) Delete(key interface{}) (value interface{}, exists bool) {
	shard := cm.GetShard(key)

	shard.Lock()
	value, exists = shard.items[key]
	delete(shard.items, key)
	shard.Unlock()

	return value, exists
}

func (cm GConcMap) Len() int {
	count := 0
	for i := 0; i < int(cm.ShardCount); i++ {
		shard := cm.shards[i]
		shard.RLock()
		count += len(shard.items)
		shard.RUnlock()
	}
	return count
}

func (cm GConcMap) GetShard(key interface{}) *GMapShard {
	return cm.shards[cm.Resolver(key, cm.ShardCount)]
}
